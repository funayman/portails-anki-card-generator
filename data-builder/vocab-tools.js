
// ensure you have all the correct sections selected

var arr = []
var entries = $('.c-row.c-row--vocab-tools')

for(i=1; i<entries.length; i++) {
  var that = entries[i];

	// skip headers and whatnot
  if(that.childElementCount == 1) {
    continue
  }

  // build out note fields
  var e = {
    "french": that.childNodes[0].textContent.trim(),
    "english": that.childNodes[1].textContent.trim(),
    "image": null,
    "audio": null,
    "base": that.childNodes[0].textContent.trim().match(/[a-zA-ZÀ-ÿ]+/g).join('-').toLowerCase()
  }

  arr.push(e)
}

console.log(JSON.stringify(arr))
console.log("copy object")
