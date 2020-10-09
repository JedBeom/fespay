export default safariDetect => {
    let iOS = /iPad|iPhone|iPod/.test(navigator.userAgent) && !window.MSStream;
    if (iOS) {
        let ua = window.navigator.userAgent;
        let avoidBrowsers = ["FxiOS", "NAVER", "KAKAOTALK"];
        for (let i = 0; i < avoidBrowsers.length; i++) {
            if (ua.includes(avoidBrowsers[i])) {
                return false
            }
        }
    }
    return true
}