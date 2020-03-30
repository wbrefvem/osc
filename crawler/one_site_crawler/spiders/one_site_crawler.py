from datetime import datetime
from scrapy import Spider
from scrapy import signals
from urllib.parse import urlparse


class OSCSpider(Spider):
    name = "osc"

    def __init__(self, start_urls=None, allowed_domains=None, *args, **kwargs):
        if start_urls:
            self.start_urls = start_urls.split(',')
        
        if allowed_domains:
            domain_list = allowed_domains.split(',')
            parsed_url = urlparse(domain_list[0])
            if parsed_url.hostname:
                self.allowed_domains = [parsed_url.hostname]
            else:
                self.allowed_domains = [domain_list[0]]
        
        super(OSCSpider, self).__init__(*args, **kwargs)

    def parse(self, response):
        yield {
            'url': response.url,
            'last-accessed': datetime.utcnow().isoformat(),
        }
        for link in response.css('a'):
            try:
                yield response.follow(link, callback=self.parse)
            except Exception as err:
                pass
