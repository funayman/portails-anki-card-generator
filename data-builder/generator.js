/*
{
  "french": "that.childNodes[3].childNodes[5].childNodes[5].dataset.title",
  "english": "that.childNodes[3].childNodes[5].childNodes[5].dataset.translation",
  "image": "that.childNodes[1].childNodes[1].src",
  "audio": "that.dataset.audioPath"
}
*/

var arr = []
var entries = $('.js-group-entry')

for(i=0; i<entries.length; i++) {
  var that = entries[i];
  var e = {
    "french": that.childNodes[3].childNodes[5].childNodes[5].dataset.title,
    "english": that.childNodes[3].childNodes[5].childNodes[5].dataset.translation,
    "image": that.childNodes[1].childNodes[1].src,
    "audio": that.dataset.audioPath,
    "base": that.childNodes[3].childNodes[5].childNodes[5].dataset.title.match(/[a-zA-ZÀ-ÿ]+/g).join('-').toLowerCase()
  }
  arr.push(e)
}

// copy JSON to clipboard

// print to console
