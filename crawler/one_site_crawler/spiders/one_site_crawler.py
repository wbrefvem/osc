import scrapy
from datetime import datetime


class OSCSpider(scrapy.Spider):
    name = "osc"

    def __init__(self, start_urls=None, allowed_domains=None, *args, **kwargs):
        if start_urls:
            self.start_urls = start_urls.split(',')
        
        if allowed_domains:
            self.allowed_domains = allowed_domains.split(',')
        
        super(OSCSpider, self).__init__(*args, **kwargs)


    def parse(self, response):
        yield {
            'url': response.url,
            'timestamp': datetime.utcnow(),
        }
        for link in response.css('a'):
            try:
                yield response.follow(link, callback=self.parse)
            except Exception as err:
                pass
