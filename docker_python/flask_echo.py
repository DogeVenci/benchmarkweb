from flask import Flask, jsonify
import json
import pymongo
from pymongo import MongoClient
client = MongoClient('mongodb://root:example@mongo:27017/')
db = client['test']

app = Flask(__name__)


@app.route('/')
def hello_world():
    # results = db.items.find_one({"filename": "28477.m3u8"}, {"_id": 0})
    results = []
    for item in db.items.find({}, {"_id": 0}):
        results.append(item)

    # return json.dumps(results)
    return jsonify(results)


if __name__ == '__main__':
    app.run(debug=True)
