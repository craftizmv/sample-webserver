# blive-coding-exercise


## Task 1

Write a basic HTTP web server which simply returns 200 OK on requests to `/healthcheck`


## Task 2

Extend this server to return metadata about live channels from files stored on disk.

Write an endpoint `/v1/channels/{channel_id}` which returns json stored in `data/channels/<channel_id>.json`. If the file does not exist, return a 404.

## Task 3

Extend this server to enrich the metadata with additional information from the Blive API.

Write an endpoint `/v2/channels/{channel_id}` which returns json stored in `data/channels/<channel_id>.json` with the ingest point and playback point added based on the Blive API.
- If the file does not exist, return a 404.
- If the Blive API returns an error, return a 500.