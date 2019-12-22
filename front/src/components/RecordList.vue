<template>
    <div>
        <table>
            <thead>
                <tr>
                    <td>ID</td>
                    <td>유형</td>
                    <td>고객</td>
                    <td>금액</td>
                    <td>결제 일시</td>
                    <td>취소 일시</td>
                </tr>
            </thead>
            <tbody>
                <tr v-for="r in records" v-bind:key="r.id">
                    <td>{{r.id}}</td>
                    <td>{{r.type}}</td>
                    <td>{{r.user.name}}</td>
                    <td>{{r.amount}}</td>
                    <td>{{r.hours}}:{{r.minutes}}:{{r.seconds}}</td>
                    <template v-if="r.canceledAt">
                    <td>{{r.canceledAt.hours}}:{{r.canceledAt.minutes}}:{{r.canceledAt.seconds}}</td>
                    </template>
                    <template v-else>
                    <td>-</td>
                    </template>
                </tr>
            </tbody>
        </table>
    </div>    
</template>

<script>
export default {
    name: "Records",   
    props: ["records"],
}
</script>

<style scoped>
    
</style>