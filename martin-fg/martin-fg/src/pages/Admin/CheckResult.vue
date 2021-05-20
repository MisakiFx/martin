<template>
  <div class="check-result">
    <div class="check-head">
      体检结果登记
    </div>
    <div class="check-field">
      <mt-field label="手机号" placeholder="请输入用户手机号" type="tel" v-model="phone"></mt-field>
      <mt-radio
        title="请选择体检项目"
        v-model="finishProject"
        :options="finishProjectOptions">
      </mt-radio>
      <div class="field">
        <mt-field label="体检结果" placeholder="请输入结果" type="textarea" v-model="result" :readonly=false rows="4"></mt-field>
      </div>
    </div>
    <button @click="checkResultButton">提交</button>
  </div>
</template>

<script>
  import {MessageBox, Toast} from "mint-ui";
  import {checkResult} from "../../api";

  export default {
      name: "CheckResult",
      data() {
        return {
          phone : '',
          finishProject : '',
          result : '',
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
        checkResultButton () {
          MessageBox.confirm("是否提交?").then(async action=>{
            let result = await checkResult({
              'phone_number' : this.phone,
              'check_project' : +this.finishProject,
              'check_result' : this.result
            }, this.$store.state.userOpenId)
            if (result.code !== 0) {
              Toast(result.msg)
              return
            }
            Toast("体检结果登记成功")
            setTimeout(()=>{
              this.$router.replace('/admin/top')
            }, 1000)
          })
        }
      }
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .check-result
    width  100%
    height 100%
    .check-head
      width 100%
      height 30px
      line-height 50px
      font-weight bold
      text-align center
    .check-field
      margin-top 25px
      .field
        line-height 50px
    button
      width 90%
      height 40px
      line-height 40px
      background-color: #e9232c
      text-align center
      margin-top 20px
      margin-left 5%
      margin-right 5%
      border none
      font-size 16px
      color #fff
      border-radius 10px
</style>
