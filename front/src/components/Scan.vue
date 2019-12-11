<template>
    <div>
        <h2>{{message}}</h2>
        <video id="video"></video> 
    </div>
</template>

<script>
import { BrowserBarcodeReader } from "@zxing/library"

export default {
    data: function() {
        return {code: "", message: "스캔 준비 중", 
        width: window.screen.availWidth, height: window.screen.availHeight}
    },
    methods: {
        ready: function() {
            if (this.message === "로딩 중...") {
                return
            }
            this.message = "스캔 중"
            const codeReader = new BrowserBarcodeReader();

            codeReader
            .decodeFromInputVideoDevice(undefined, 'video')
            .then(result => {
                codeReader.reset();
                this.message = "로딩 중..."
                this.$emit("on-detect", result.text);
            })
            /* eslint-disable no-unused-vars */
            .catch(err => {
                this.$emit("on-detect", "");
                alert("카메라를 실행할 수 없습니다. 카메라 권한을 허용했는지 확인해주세요.");
            });
            /* eslint-enable no-unused-vars */
        }
    },
    mounted: function () {
        this.$nextTick(function () {
            this.ready();
        })
    }
}
</script>

<style scoped>
h2 {
    margin: 20px;
    color: white;
    text-shadow: 2px 2px 2px black;
    margin-top: 20vh;
    position: fixed;
}

video {
    top: 0;
    left: 0;
    z-index: -100;
    position: fixed;
}
</style>