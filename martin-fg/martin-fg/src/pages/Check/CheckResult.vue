<template>
    <div class="check-result-container">
      <ul class="check-list">
        <li class="check-list-item" v-show="checkResult.projects.indexOf(1) > -1">
          <div class="list-item-head">
            <h4 class="item-time-label">体检项目</h4>
            <h4 class="item-time-value">内科</h4>
          </div>
          <div class="list-item-bottom">
            <mt-field  class="item-result" label="体检结果" type="textarea" v-model="checkResult.internal" :readonly=true :disabled=true rows="4"></mt-field>
          </div>
        </li>
        <li class="check-list-item" v-show="checkResult.projects.indexOf(2) > -1">
          <div class="list-item-head">
            <h4 class="item-time-label">体检项目</h4>
            <h4 class="item-time-value">外科</h4>
          </div>
          <div class="list-item-bottom">
            <mt-field class="item-result" label="体检结果" type="textarea" v-model="checkResult.surgery" :readonly=true :disabled=true rows="4"></mt-field>
          </div>
        </li>
        <li class="check-list-item" v-show="checkResult.projects.indexOf(3) > -1">
          <div class="list-item-head">
            <h4 class="item-time-label">体检项目</h4>
            <h4 class="item-time-value">耳鼻喉科</h4>
          </div>
          <div class="list-item-bottom">
            <mt-field class="item-result" label="体检结果" type="textarea" v-model="checkResult.ent" :readonly=true :disabled=true rows="4"></mt-field>
          </div>
        </li>
        <li class="check-list-item" v-show="checkResult.projects.indexOf(4) > -1">
          <div class="list-item-head">
            <h4 class="item-time-label">体检项目</h4>
            <h4 class="item-time-value">肝功</h4>
          </div>
          <div class="list-item-bottom">
            <mt-field class="item-result" label="体检结果" type="textarea" v-model="checkResult.sgpt" :readonly=true :disabled=true rows="4"></mt-field>
          </div>
        </li>
        <li class="check-list-item" v-show="checkResult.projects.indexOf(5) > -1">
          <div class="list-item-head">
            <h4 class="item-time-label">体检项目</h4>
            <h4 class="item-time-value">血糖</h4>
          </div>
          <div class="list-item-bottom">
            <mt-field class="item-result" label="体检结果" type="textarea" v-model="checkResult.blood_glucode" :readonly=true :disabled=true rows="4"></mt-field>
          </div>
        </li>
        <li class="check-list-item" v-show="checkResult.projects.indexOf(6) > -1">
          <div class="list-item-head">
            <h4 class="item-time-label">体检项目</h4>
            <h4 class="item-time-value">血脂</h4>
          </div>
          <div class="list-item-bottom">
            <mt-field class="item-result" label="体检结果" type="textarea" v-model="checkResult.blood_fat" :readonly=true :disabled=true rows="4"></mt-field>
          </div>
        </li>
        <li class="check-list-item" v-show="checkResult.projects.indexOf(7) > -1">
          <div class="list-item-head">
            <h4 class="item-time-label">体检项目</h4>
            <h4 class="item-time-value">肾功</h4>
          </div>
          <div class="list-item-bottom">
            <mt-field class="item-result" label="体检结果" type="textarea" v-model="checkResult.renal_function" :readonly=true :disabled=true rows="4"></mt-field>
          </div>
        </li>
      </ul>
    </div>
</template>

<script>
    import {getCheckResult} from '../../api'
    import {Toast} from 'mint-ui'
    export default {
      name: "CheckResult",
      mounted() {
        this.init()
      },
      methods : {
        async init() {
          let result = await getCheckResult(this.$store.state.userInfo.open_id, this.$route.params.id)
          if (result.code !== 0) {
            Toast(result.msg)
            return
          }
          this.checkResult = result.data.result
          this.checkResult.internal = this.checkResult.internal ? this.checkResult.internal : '请耐心等待体检结果'
          this.checkResult.surgery = this.checkResult.surgery ? this.checkResult.surgery : '请耐心等待体检结果'
          this.checkResult.ent = this.checkResult.ent ? this.checkResult.ent : '请耐心等待体检结果'
          this.checkResult.sgpt = this.checkResult.sgpt ? this.checkResult.sgpt : '请耐心等待体检结果'
          this.checkResult.blood_glucode = this.checkResult.blood_glucode ? this.checkResult.blood_glucode : '请耐心等待体检结果'
          this.checkResult.blood_fat = this.checkResult.blood_fat ? this.checkResult.blood_fat : '请耐心等待体检结果'
          this.checkResult.renal_function = this.checkResult.renal_function ? this.checkResult.renal_function : '请耐心等待体检结果'
        }
      },
      data() {
        return{
          checkResult : {
            projects : []
          },
        }
      }
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .check-result-container
    padding-bottom 10px
    background-color: #f5f5f5
    .checks-empty
      text-align center
      color #9b9898
      padding-top 50%
      line-height 10px
      font-size 12px
    .check-list
      .check-list-item
        background-color: #fff
        display flex
        flex-direction column
        margin 10px 0px
        .list-item-head
          margin 10px 10px
          display flex
          justify-content space-around
          align-items center
          .item-time-value
            font-size 15px
            weight 50%
          .item-time-label
            line-height 20px
            width 50%
            height 20px
            color black
            overflow hidden
            font-size 15px
        .list-item-bottom
          margin 0px 10px
          margin-bottom 10px
          height 105px
          display flex
          flex-direction row
          align-items center
          font-size 13px
          color #736c6c
          position relative
          .item-result
            height 100%
            width 100%
            line-height 100px
    .expense-tip
      text-align center
      color #9b9898
      line-height 10px
      font-size 12px
</style>
