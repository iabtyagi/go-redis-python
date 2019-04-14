
import time, json
from redis import StrictRedis
from redis.exceptions import RedisError


time.sleep(8)

redisDB = StrictRedis(host='redis', port=6379, db=0, password='abcd1234', decode_responses=True)

payload = {
            "userid": 1234,
            "message": "You got a message. ' \" \ @ # & % 😀",
            "message_type": 2
            }

while True:
    payload["ts"] = int(time.time())
    # Right push to the list. We'll pop from left.
    redisDB.rpush('gopy_message_queue', json.dumps(payload))
    time.sleep(5)
