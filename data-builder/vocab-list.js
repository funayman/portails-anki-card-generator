
var arr = []
var entries = $('.c-row')

for(i=0; i<entries.length; i++) {
  var that = entries[i];
  var e = {
    "french": that.childNodes[1].textContent.trim(),
    "english": that.childNodes[3].textContent.trim(),
    "image": null,
    "audio": that.childNodes[1].childNodes[1].childNodes[1].currentSrc,
    "base": that.childNodes[1].textContent.trim().match(/[a-zA-ZÀ-ÿ]+/g).join('-').toLowerCase()
  }
  arr.push(e)
}

console.log(JSON.stringify(arr))
console.log("copy object")
