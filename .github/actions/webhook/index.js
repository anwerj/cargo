const fs = require('fs');
const https = require('https');

const apiEndpoint = 'https://enmj69739hb8l.x.pipedream.net/push/github';

function push(data) {
    try {
        const response = https.request(apiEndpoint, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
        }, (res) => {
            const statusCode = res.statusCode;
            if (statusCode === 200) {
                console.log('Successfully sent events to collector');
            } else {
                console.log(`Non 200 status code:${statusCode} received from collector endpoint`);
            }
            process.exit(0);
        });

        response.on('error', (error) => {
            console.error('Exception:', error.message);
            // Don't want to block pipelines because of failure in this script
            process.exit(0);
        });

        response.write(JSON.stringify(data));
        response.end();
    } catch (error) {
        console.error('Exception:', error.message);
        // Don't want to block pipelines because of failure in this script
        process.exit(0);
    }
}

function removeKeysWithSuffix(data, suffixes) {
    if (typeof data !== 'object') {
        return data
    }
    if (Array.isArray(data)) {
        return data.map(item => removeKeysWithSuffix(item, suffixes));
    } else {
        const newData = {};
        for (const key in data) {
            if (!suffixes.some(suffix => key.endsWith(suffix))) {
                newData[key] = removeKeysWithSuffix(data[key], suffixes);
            }
        }
        return newData;
    }
}

async function main() {
    const eventFile = process.env.GITHUB_EVENT_PATH;
    try {
        const eventData = fs.readFileSync(eventFile, 'utf-8');
        const data = JSON.parse(eventData);
        const filteredData = removeKeysWithSuffix(data, ['_url', '_links', 'repo', 'url']);
        // Update the watchtower with the same content as the file has
        push(filteredData);
    } catch (error) {
        console.error('Exception:', error.message);
        // Don't want to block pipelines because of failure in this script
        process.exit(0);
    }
}

main();
