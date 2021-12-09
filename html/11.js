function ran() {
    return parseInt(Math.random() * 256)
}
console.log(ran());

var json = '{"name":"黄雪朋", "age":12}'
json1 = JSON.parse(json)
console.log(json1.name);