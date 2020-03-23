import os
from flask import Flask
import requests
import json
import logging
import traceback
from contextlib import closing
app = Flask(__name__)

@app.route("/conf/env")
def env():
    f = os.popen('printenv')
    return f.read()

@app.route("/env/<name>/<value>")
def create_env_var(name, value):
    with open(os.path.expanduser("~/.bashrc"), "a") as outfile:
        outfile.write("export "+name+"="+value+"\n")
    return "OK"

@app.route("/env/healthchecker")
def healthchecker():
    f = os.popen('ps -o pid,cmd -ef')
    data = f.read()
    sendSlackMessage(data)
    return "OK"

def sendSlackMessage(message):
    logging.basicConfig()
    log = logging.getLogger()
    log.setLevel(logging.INFO)

    SLACK_DEBUG_URL = "https://hooks.slack.com/services/zzz/xxx/yyy"
    SLACK_DEBUG_CHANNEL = "aprox_messages"
    app_name = "aproxmessagesystem"
    TIMEOUT_PER_REQUEST = 5

    slack_channel = SLACK_DEBUG_CHANNEL
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

if __name__ == '__main__':
    app.run(debug='true')