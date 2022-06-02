let x = 2;
puts(x);

let y = [1,2,3];
puts(y);

let addTwo = fn(x) { 
    x + 2 
};

let y = y.map(addTwo);
puts(y);

let person = { "name": "Tom Bombadil", "clothes": { "shoes": "yellow boots" } };
let shoes = person.dig("clothes", "shoes");
puts(shoes);

if (1 > 2) {
    puts("a")
} else {
    puts("b")
};

for(let x = 0; x < 10; x = x + 1) {
    puts(x)
};