# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html
import os
import json

from urllib.parse import urlparse
from os import path
from datetime import datetime

class OneSiteCrawlerPipeline(object):

    def process_item(self, item, spider):
        self.outfile.seek(0)
        try:
            sitemap = json.load(self.outfile)
        except json.JSONDecodeError:
            self.outfile.close()
            spider.crawler.stop()
        sitemap['data'].append(item)

        self.outfile.seek(0)
        self.outfile.truncate()
        try:
            json.dump(sitemap, self.outfile)
        except json.JSONDecodeError:
            self.outfile.close()
            spider.crawler.stop()
        return item

    def open_spider(self, spider):
        data_dir = os.getenv("DATA_DIR", "/opt/data")
        domain = spider.allowed_domains[0]
        file_path = path.join(data_dir, "%s.json" % domain)

        sitemap = {
            'last-crawled': 'never',
            'domain': domain,
            'crawl-status': 'pending',
            'data': [],
        }

        self.outfile = open(file_path, 'a+')
        self.outfile.seek(0)
            
        # If we just created the file, the JSON decode will fail
        try:
            data = json.load(self.outfile)
            self.outfile.seek(0)
        except json.JSONDecodeError:
            data = sitemap
        
        self.outfile.truncate()
        json.dump(data, self.outfile)
    
    def close_spider(self, spider):
        self.outfile.seek(0)
        try:
            data = json.load(self.outfile)
        except json.JSONDecodeError:
            spider.crawler.stop()
        data['crawl-status'] = 'complete'
        data['last-crawled'] = datetime.utcnow().isoformat()

        self.outfile.seek(0)
        self.outfile.truncate()
        try:
            json.dump(data, self.outfile)
        except json.JSONDecodeError:
            spider.crawler.stop()
        self.outfile.close()
