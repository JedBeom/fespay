<template>
    <video id="video"></video> 
</template>

<script>
import { BrowserBarcodeReader } from "@zxing/library"

export default {
    data: function() {
        return {code: "", 
        width: window.screen.availWidth, height: window.screen.availHeight, codeReader: BrowserBarcodeReader()}
    },
    props: {
        active: Boolean
    },
    methods: {
        ready: function(loader) {
            this.codeReader = new BrowserBarcodeReader();

            this.codeReader
            .decodeFromInputVideoDevice(undefined, 'video')
            .then(result => {
                this.codeReader.reset();
                this.$emit("on-detect", result.text);
            })
            /* eslint-disable no-unused-vars */
            .catch(err => {
                this.$emit("on-detect", "");
                alert("카메라를 실행할 수 없습니다. 카메라 권한을 허용했는지 확인해주세요.");
            });
            /* eslint-enable no-unused-vars */
            loader.hide()
        }
    },
    mounted () {
        let loader = this.$loading.show()
        this.$nextTick(function () {
            this.ready(loader);
        })
    },
    computed: {
        activeOrNot() {
            if (!this.active) {
                this.codeReader.reset()
            }
            return this.active
        }
    }
}
</script>

<style scoped>

</style>