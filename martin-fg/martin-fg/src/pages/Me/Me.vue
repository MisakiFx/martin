<template>
  <div class="me">
    <select-login />
  </div>
</template>

<script>
    import SelectLogin from "../Login/SelectLogin";
    export default {
        components:{
            SelectLogin
        },
        name: "Me"
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .me
    width 100%
    height 100%
    background #f5f5f5
</style>
