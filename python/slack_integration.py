#python3
import requests
import json
import logging
import traceback
from contextlib import closing

logging.basicConfig()
log = logging.getLogger()
log.setLevel(logging.INFO)

SLACK_DEBUG_URL = "https://hooks.slack.com/services/xxx/yyy/zzz"
SLACK_DEBUG_CHANNEL = "aprox_messages"
SLACK_AUTHOR = "AproxSniffer"
app_name = "aproxmessagesystem"
TIMEOUT_PER_REQUEST = 5

slack_channel = SLACK_DEBUG_CHANNEL
message = "@here THE WORLD IS GOING DOWN"
author_name = "AproxSniffer"
slack_url = SLACK_DEBUG_URL

try:
    slack_headers = {'Content-type': 'application/json', 'Accept': 'application/json', 'User-Agent': 'custom agent'}
    color = "#f90c0c"
    slackData = \
        {
            "channel": slack_channel,
            "username": app_name,
            "attachments":
                [
                    {
                        "author_name": 'Notification',
                        "title": author_name,
                        "color": color,
                        "text": message,
                        "mrkdwn_in": ["text", "pretext"]
                    }
                ]
        }

    with closing(requests.post(slack_url, data=json.dumps(slackData), headers=slack_headers, timeout=TIMEOUT_PER_REQUEST)) as response:
        the_page = response.text
        log.debug("Slack response: '%s'", the_page)
except:
    log.error("Fail to send slack message, url %s, channel: %s, message %s\n\n%s", slack_url, slack_channel, message, traceback.format_exc())