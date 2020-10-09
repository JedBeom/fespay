export default function(amount) {
    amount = amount.replace(/\D/g, '').replace(/^0+/, '')
    if (amount === "0" || amount === "" || amount === 0 || isNaN(amount) || parseInt(amount) <= 0) {
        return "금액을 입력하세요"
    }

    if (amount > 50000) {
        return "50000 코인을 초과할 수 없습니다"
    }

    if (amount % 100 !== 0) {
        return "100 코인 단위로 입력해야 합니다"
    }

    return ""
}