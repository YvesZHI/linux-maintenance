# Data structure alignment


### rule ###
1) The offset of the first element equals to `0`; the offset of the nth element must equal to the next integer of multiple of `min(sizeof(nth element), #pragma pack())`;<br>
2) After the rule 1, the padding action may be necessary to make sure that the size of the struct must equal to the next integer of multiple of `min(max(sizeof(all elements)), #pragma pack())`;<br>
3) the offset of a nested struct must equal to the next integer of multiple of `max(sizeof(all elements in the nested struct))`.

### examples ###
```
#pragma pack(1)
struct AA {
    int a;   // sizeof(a) = 4, the offset of the first element always equals to 0, so it is put at [0,3]
    char b;  // min(sizeof(b), pack()) = min(1, 1) = 1, 4 % 1 = 0, offset = 4, so it is put at [4]
    short c; // min(sizeof(c), pack()) = min(2, 1) = 1, 5 % 1 = 0, offset = 5, so it is put at [5, 6]
    char d;  // min(sizeof(d), pack()) = min(1, 1) = 1, 7 % 1 = 0, offset = 7, so it is put at [7]
};
// 8 % min(max(sizeof(all elements)), pack()) = 8 % 1 = 0
// so sizeof(AA) = 8
```

```
#pragma pack(2)
struct AA {
    int a;   // sizeof(a) = 4, the offset of the first element always equals to 0, so it is put at [0,3]
    char b;  // min(sizeof(b), pack()) = min(1, 2) = 1, 4 % 1 = 0, offset = 4, so it is put at [4]
    short c; // min(sizeof(c), pack()) = min(2, 2) = 2, 6 % 2 = 0, offset = 6, so it is put at [6, 7]
    char d;  // min(sizeof(d), pack()) = min(1, 2) = 1, 8 % 1 = 0, offset = 8, so it is put at [8]
};
// 9 % min(max(sizeof(all elements)), pack()) = 9 % 2 != 0
// so padding 1 more byte to make (9 + 1) % 2 = 0
// so sizeof(AA) = 10
```

```
#pragma pack(4)
struct AA {
    int a;   // sizeof(a) = 4, the offset of the first element always equals to 0, so it is put at [0,3]
    char b;  // min(sizeof(b), pack()) = min(1, 4) = 1, 4 % 1 = 0, offset = 4, so it is put at [4]
    short c; // min(sizeof(c), pack()) = min(2, 4) = 2, 6 % 2 = 0, offset = 6, so it is put at [6, 7]
    char d;  // min(sizeof(d), pack()) = min(1, 4) = 1, 8 % 1 = 0, offset = 8, so it is put at [8]
};
// 9 % min(max(sizeof(all elements)), pack()) = 9 % 4 != 0
// so padding 3 more bytes to make (9 + 3) % 4 = 0
// so sizeof(AA) = 12
```

```
#pragma pack(8)
struct AA {
    int a;   // sizeof(a) = 4, the offset of the first element always equals to 0, so it is put at [0,3]
    char b;  // min(sizeof(b), pack()) = min(1, 8) = 1, 4 % 1 = 0, offset = 4, so it is put at [4]
    short c; // min(sizeof(c), pack()) = min(2, 8) = 2, 6 % 2 = 0, offset = 6, so it is put at [6, 7]
    char d;  // min(sizeof(d), pack()) = min(1, 8) = 1, 8 % 1 = 0, offset = 8, so it is put at [8]
};
// 9 % min(max(sizeof(all elements)), pack()) = 9 % 4 != 0
// so padding 3 more bytes to make (9 + 3) % 4 = 0
// so sizeof(AA) = 12
```

```
#pragma pack(8)
struct EE
{
    int a;      // sizeof(a) = 4, the offset of the first element always equals to 0, so it is put at [0,3]
    char b;     // min(sizeof(b), pack()) = min(1, 8) = 1, 4 % 1 = 0, offset = 4, so it is put at [4]
    short c;    // min(sizeof(c), pack()) = min(2, 8) = 2, 6 % 2 = 0, offset = 6, so it is put at [6, 7]
    // max(sizeof(all elements of FF)) = 4, 8 % 4 = 0, so it is put at [8, ?]
    struct FF
    {
        int a1;     // min(sizeof(a1), pack()) = min(4, 8) = 4, 8 % 4 = 0, offset = 8, so it is put at [8, 11]
        char b1;    // min(sizeof(b1), pack()) = min(1, 8) = 1, 12 % 1 = 0, offset = 12, so it is put at [12]
        short c1;   // min(sizeof(c1), pack()) = min(2, 8) = 2, 14 % 2 = 0, offset = 14, so it is put at [14, 15]
        char d1;    // min(sizeof(d1), pack()) = min(1, 8) = 1, 16 % 1 = 0, offset = 16, so it is put at [16]
    };
    // 17 % min(max(sizeof(all elements)), pack()) = 17 % 4 != 0
    // so padding 3 more bytes to make (17 + 3) % 4 = 0
    char d;         // min(sizeof(d), pack()) = min(1, 8) = 1, 21 % 1 = 0, offset = 21, so it is put at [21]
};
// 22 % min(max(sizeof(all elements)), pack()) = 22 % 4 != 0
// so padding 2 more bytes to make (22 + 2) % 4 = 0
// so sizeof(EE) = 24
```

```
#pragma pack(8)
struct B {
    char e[2];      // sizeof(e) = 2, the offset of the first element always equals to 0, so it is put at [0,1]
    short h;        // min(sizeof(h), pack()) = min(2, 8) = 2, 2 % 2 = 0, offset = 2, so it is put at [2, 3]
    // max(sizeof(all elements of A)) = 8, 8 % 8 = 0, so it is put at [8, ?]
    struct A {
        int a;      // min(sizeof(a), pack()) = min(4, 8) = 4, 8 % 4 = 0, offset = 8, so it is put at [8, 11]
        double b;   // min(sizeof(b), pack()) = min(8, 8) = 8, 16 % 8 = 0, offset = 16, so it is put at [16, 23]
        float c;    // min(sizeof(c), pack()) = min(4, 8) = 4, 24 % 4 = 0, offset = 24, so it is put at [24, 27]
    };
    // 28 % min(max(sizeof(all elements)), pack()) = 28 % 8 != 0
    // so padding 4 more bytes to make (28 + 4) % 8 = 0
};
// 32 % min(max(sizeof(all elements)), pack()) = 32 % 8 = 0
// so sizeof(B) = 32
