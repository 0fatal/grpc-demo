import bs4
import requests as req
from enum import Enum
from fake_useragent import UserAgent

class NewsType(Enum):
    ALL = 0
    GAME=229
    PLAYER=231
    EQUIPMENT=232
    FUN=230

class SectionType(Enum):
    TEACH = 36
    NEWS = 239
    EQUIPMENT = 227
    GAME_VIDEO = 2
    ASK = 329

def fetchPPW(page=1,section_type = SectionType.NEWS, news_type = NewsType.ALL):
    ua = UserAgent()
    headers = {
        'User-Agent': ua.random
    }
    url = ''
    if news_type.value == 0:
        url = f'https://www.pingpangwang.com/forum.php?mod=forumdisplay&fid={section_type.value}&page={page}'
    else:
        url =  f'https://www.pingpangwang.com/forum.php?mod=forumdisplay&fid={section_type.value}&filter=typeid&typeid={news_type.value}&page={page}'

    resBody = req.get(url,headers=headers)
    text = resBody.content.decode(resBody.apparent_encoding)
    res = []

    nodes = bs4.BeautifulSoup(text,'html.parser').find_all('tbody',id=lambda s:s.startswith('normalthread_'))
    for node in nodes:
        title = node.tr.select_one('th > a[class="s xst"]').string
        link = node.tr.select_one('th > a[class="s xst"]')['href']
        try:
            time = node.tr.select_one('.by > em > span > span').string
            date = node.tr.select_one('.by > em > span > span')['title']
        except AttributeError:
            time = date = node.tr.select_one('.by > em > a').string
        res.append({
            'title': title,
            'link': link,
            'time': time,
            'date': date
        })
    return res

# print(fetchPP(section_type=SectionType.GAME_VIDEO))