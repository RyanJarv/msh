function handler () {
    EVENT_DATA=$1

    msg="$(echo -n "$EVENT_DATA"|jq -r ".[0].msg")"
    event="$(echo -n "$EVENT_DATA"|jq -r ".[0].event")"

    if [[ $event == "ClosedEvent" ]]; then
      pipeName="$(echo -n "$EVENT_DATA"|jq -r ".[0].pipeName")"
      region="$(echo -n "$EVENT_DATA"|jq -r ".[0].awsRegion")"

      if [[ $pipeName == "null" ]]; then
        echo 1>&2 "[ERROR] No .pipeName key found in ClosedEvent event"
        echo "{}"; return
      fi

      if [[ $region == "null" ]]; then
        echo 1>&2 "[ERROR] No .awsRegion key found in ClosedEvent event"
        echo "{}"; return
      fi

      code=0
      aws --region "$region" pipes stop-pipe --name "$pipeName" || code=$?

      if [[ $code -ne 0 ]]; then
        echo 1>&2 "[ERROR] stop-pipe call failed on '$pipeName'"
        echo '{}'; return
      fi
    fi
  resp="$(main $msg)"
  echo '{}' | jq --arg resp "$resp" '{"statusCode": 200, "body": $resp}'

}

function main() {
  MSG=$1
  RESPONSE="Hello from Lambda!"
  echo $RESPONSE
}
