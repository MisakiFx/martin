<template>
  <div class="check-finish">
    <div class="check-head">
      体检登记
    </div>
    <div class="check-field">
      <mt-field label="手机号" placeholder="请输入用户手机号" type="tel" v-model="phone"></mt-field>
      <mt-radio
        title="请选择体检日期(仅支持未来三天内预约)"
        v-model="finishProject"
        :options="finishProjectOptions">
      </mt-radio>
    </div>
    <button @click="checkFinishButton">提交</button>
  </div>
</template>

<script>
  import {MessageBox, Toast} from 'mint-ui'
  import {checkFinish} from "../../api";

  export default {
      name: "CheckFinish",
      data() {
        return {
          phone : '',
          finishProject : '',
          finishProjectOptions:[
            {
              label: '内科',
              value: '1'
            },
            {
              label: '外科',
              value: '2'
            },
            {
              label: '耳鼻喉科',
              value: '3'
            },
            {
              label: '肝功',
              value: '4'
            },
            {
              label: '血糖',
              value: '5'
            },
            {
              label: '血脂',
              value: '6'
            },
            {
              label: '肾功',
              value: '7'
            },
          ],
        }
      },
      methods : {
        checkFinishButton() {
          MessageBox.confirm("是否提交?").then(async action=>{
            let result = await checkFinish({
              'phone_number' : this.phone,
              'finish_project' : +this.finishProject
            }, this.$store.state.userOpenId)
            if (result.code !== 0) {
              Toast(result.msg)
              return
            }
            Toast("单项体检结束")
            setTimeout(()=>{
              this.$router.replace('/admin/top')
            }, 1000)
          })
        }
      }
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .check-finish
    width  100%
    height 100%
    .check-head
      width 100%
      height 50px
      line-height 50px
      font-weight bold
      text-align center
    .check-field
      margin-top 25px
    button
      width 90%
      height 40px
      line-height 40px
      background-color: #e9232c
      text-align center
      margin 50px 5%
      border none
      font-size 16px
      color #fff
      border-radius 10px
</style>
