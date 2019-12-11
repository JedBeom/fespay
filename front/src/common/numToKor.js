export function numToKor(num) {
    let hanA = ["", "일", "이", "삼", "사", "오", "육", "칠", "팔", "구", "십"];
    let danA = ["", "십", "백", "천", "", "십", "백", "천", "", "십", "백", "천", "", "십", "백", "천"];
    let result = "";
    for (let i = 0; i < num.length; i++) {
        let str = "";
        let han = hanA[num.charAt(num.length - (i + 1))];
        if (han !== "") str += han + danA[i];
        if (i === 4) str += "만";
        result = str + result;
    }
    return result;
}