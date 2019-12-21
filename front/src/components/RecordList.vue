<template>
    <div>
        <table>
            <thead>
                <tr>
                    <td>ID</td>
                    <td>고객</td>
                    <td>금액</td>
                    <td>일시</td>
                </tr>
            </thead>
            <tbody>
                <tr v-for="r in records" v-bind:key="r.id">
                    <td>{{r.id}}</td>
                    <td>{{r.user.name}}</td>
                    <td>{{r.amount}}</td>
                    <td>{{r.hours}}:{{r.minutes}}:{{r.seconds}}</td>
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