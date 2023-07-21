import sys
import re
import requests
import os
import json
from datetime import datetime
from dateutil import tz

api_endpoint = "https://enmj69739hb8l.x.pipedream.net/push/pr_merge"

def push(data):
  try:
    r = requests.post(api_endpoint, data=json.dumps(data), headers={'Content-type': 'application/json'})
    status_code = r.status_code
    if status_code == 200:
      print ("Successfully sent events to watchtower")
    else:
      print ("Non 200 status code : {}, received from watchtower endpoint".format(status_code))
  except Exception as e:
    print ("Exception: " + str(e))


def remove_keys_with_suffix(data, suffixes):
    if isinstance(data, dict):
        return {
            key: remove_keys_with_suffix(value, suffixes)
            for key, value in data.items()
            if not any(key.endswith(suffix) for suffix in suffixes)
        }
    elif isinstance(data, list):
        return [remove_keys_with_suffix(item, suffixes) for item in data]
    else:
        return data

def main():
  event_file = os.getenv("GITHUB_EVENT_PATH")
  try:
    fd = open(event_file, "r")

  except Exception as e:
    print ("Exception: " +  str(e))
    ## Don't want to block pipelines because of failure in this script
    sys.exit(0)

  data = remove_keys_with_suffix(json.load(fd),  ["_url", "_links", "url", "repository", "repo", "sender"])

  # Update the watchtower with the same content as the file has
  push(data)


if __name__ == "__main__":
    main()
