## ES 索引相关

```bash
put /object
{
    "settings": {
        "number_of_shards":1,
        "number_of_replicas":2,
        "analysis" : {
            "analyzer" : {
                "ik" : {
                    "tokenizer" : "ik_max_word"
                }
            }
    }
   },
   "mappings":{
     "properties":{
        "name": {"type": "text", "analyzer": "ik_max_word","search_analyzer": "ik_smart"},
        "tags": {"type": "text", "analyzer": "ik_max_word","search_analyzer": "ik_smart"},
        "content": {"type": "text", "analyzer": "ik_max_word","search_analyzer": "ik_smart"},
        "node_id": {"type": "long"},
        "app": {"type": "keyword"}
      }
   }
}
```


