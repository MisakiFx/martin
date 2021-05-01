<template>
  <div class="check-list-container">
    <div class="checks-empty" v-show="checks.length === 0">无体检信息</div>
    <ul class="check-list"
        v-infinite-scroll="scrollMore"
        infinite-scroll-disabled="moreShowBool"
        infinite-scroll-distance="0">
      <li class="check-list-item" v-for="(check, index) in checks" :key="index">
        <router-link tag="div" class="check-list-item" :to="'/check_result/' + check.id">
          <div class="list-item-head">
            <h4 class="item-time-label">预约时间</h4>
            <h4 class="item-time-value">{{check.start_time}}</h4>
          </div>
          <div class="list-item-bottom">
            <span class="item-project">{{check.check_project | projectFilter}}</span>
            <span class="item-money">付款:{{check.pay_check_count !== 0 ? '体检套餐' : '¥' + check.pay_reminder}}</span>
            <button class="item-button" v-show="check.status === 0" @click="cancelCheck(check.id)">取消预约</button>
          </div>
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script>
    import {getUserCheckList, postCancelCheck} from "../../api";
    import {MessageBox, Toast} from "mint-ui";

    export default {
      name: "CheckList",
      data(){
        return {
          checks : [],
          moreShowBool : false,
          allLoaded : false,
          page : 0,
        }
      },
      filters : {
        projectFilter(projectInts) {
          let resProject = []
          for (let i in projectInts) {
            switch (projectInts[i]) {
              case 1 :
                resProject.push('内科')
                break
              case 2 :
                resProject.push('外科')
                break
              case 3 :
                resProject.push('耳鼻喉科')
                break
              case 4 :
                resProject.push('肝功')
                break
              case 5 :
                resProject.push('血糖')
                break
              case 6 :
                resProject.push('血脂')
                break
              case 7 :
                resProject.push('肾功')
                break
            }
          }
          return resProject.join(',')
        }
      },
      methods:{
        async scrollMore() {
          if (this.allLoaded === true) {
            return
          }
          this.moreShowBool = true
          this.page += 1
          let result = await getUserCheckList(this.$store.state.userInfo.open_id, this.page)
          if (result.code !== 0) {
            Toast(result.msg)
            this.moreShowBool = false
            return
          }
          this.checks = this.checks.concat(result.data.list)
          if (result.data.list.length === 0) {
            this.allLoaded = true
          }
          this.moreShowBool = false
        },
        async cancelCheck(id = 0) {
          MessageBox.confirm('确定取消预约?').then(async ()=> {
            let result = await postCancelCheck(this.$store.state.userInfo.open_id, id)
            if (result.code !== 0) {
              Toast(result.msg)
              return
            }
            Toast('取消成功')
            this.checks.shift()
          })
        }
      }
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .check-list-container
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
        margin-bottom 10px
        background-color: #fff
        display flex
        flex-direction column
        .list-item-head
          margin 10px 10px
          display flex
          justify-content space-around
          align-items center
          .item-time-value
            font-size 15px
            weight 50%
          .item-time-label
            line-height 40px
            width 50%
            height 40px
            color black
            overflow hidden
            font-size 15px
        .list-item-bottom
          margin 15px 10px
          height 15px
          display flex
          flex-direction row
          justify-content space-around
          align-items center
          font-size 13px
          color #9b9898
          position relative
          .item-project
            position absolute
            left 7px
            width 40%
          .item-money
            text-align center
            justify-content flex-end
            position absolute
            right 75px
          .item-button
            position absolute
            right 0px
            line-height 30px
            background-color: #e9232c
            text-align center
            border none
            color #fff
            border-radius 10px
    .expense-tip
      text-align center
      color #9b9898
      line-height 10px
      font-size 12px
</style>
