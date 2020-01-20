var imMem = imMem || {};
var memHash = imMem.Method = { 
    hash:new Map(),
    add:function(key, value){
        this.hash.set(key, value)
    },
    find:function(key){
       return this.hash.get(key) != null
    },
    del:function(key){
        this.hash.delete(key)
    },
    get:function(key){
        return this.hash.get(key)
    }
} 
var actRecord = actRecord || {};
var actRecords = actRecord.Method = { 
    hash:new Map(),
    add:function(key, value){
        this.hash.set(key, value)
    },
    find:function(key){
       return this.hash.get(key) != null
    },
    del:function(key){
        this.hash.delete(key)
    },
    get:function(key){
        return this.hash.get(key)
    }
} 

