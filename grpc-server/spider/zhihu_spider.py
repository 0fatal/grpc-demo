import logging

import requests as req
from fake_useragent import UserAgent


def getZhiHuHot():
    ua = UserAgent()
    headers = {
        'User-Agent': ua.random,
        'upgrade-insecure-requests': '1'
    }


    try:
        json = req.get('https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true',
                       headers=headers).json()

        res = []
        for data in json['data']:
            res.append({
                'title': data['target']['title'],
                'url': data['target']['url'],
                'created': data['target']['created'],
                'answer_count': data['target']['answer_count'],
                'follower_count': data['target']['follower_count'],
                'comment_count': data['target']['comment_count'],
                'brief': data['target']['excerpt'],
                'thumbnail': data['children'][0]['thumbnail'],
                'detail': data['detail_text']
            })
        return res
    except ValueError as e:
        logging.error('get data wrong',e)
        return []
    except Exception as e:
        logging.error(e.args)
        return []


