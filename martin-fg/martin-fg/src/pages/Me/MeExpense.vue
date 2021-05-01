<template>
    <div class="expense-container">
      <div class="expenses-empty" v-show="expenses.length === 0">无体检信息</div>
      <ul class="expense-list"
          v-infinite-scroll="scrollMore"
          infinite-scroll-disabled="moreShowBool"
          infinite-scroll-distance="0">
        <li class="expense-list-item" v-for="(expense, index) in expenses" :key="index">
          <div class="list-item-head">
            <h4 class="item-time-label">¥{{expense.money}}</h4>
            <h4 class="item-time-value">{{expense.examination_name}}</h4>
          </div>
          <div class="list-item-bottom">
            <span class="item-project">{{expense.status === 1 ? '购买' : '退款'}}</span>
            <span class="item-money">{{expense.create_time}}</span>
          </div>
        </li>
      </ul>
      <div class="expense-tip">
        <mt-spinner type="snake" :size="20" v-show="moreShowBool&&!allLoaded"></mt-spinner>
        <span v-show="allLoaded">已全部加载</span>
      </div>
    </div>
</template>

<script>
  import {getUserExpense} from '../../api'
  import {Toast} from 'mint-ui'
  export default {
    name: "MeExpense",
    methods: {
      async scrollMore(){
        if (this.allLoaded === true) {
          return
        }
        this.moreShowBool = true
        this.page += 1
        let result = await getUserExpense(this.$store.state.userInfo.open_id, this.page)
        if (result.code !== 0) {
          Toast(result.msg)
          this.moreShowBool = false
          return
        }
        this.expenses = this.expenses.concat(result.data.list)
        if (result.data.list.length === 0) {
          this.allLoaded = true
        }
        this.moreShowBool = false
      }
    },
    data(){
      return {
        expenses : [],
        moreShowBool : false,
        allLoaded : false,
        page : 0,
      }
    },
  }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .expense-container
    padding-bottom 10px
    background-color: #f5f5f5
    .expenses-empty
      text-align center
      color #9b9898
      padding-top 50%
      line-height 10px
      font-size 12px
    .expense-list
      .expense-list-item
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
            margin-right 10px
            font-size 17px
            width 20%
          .item-time-label
            line-height 40px
            width 94%
            margin-left 3%
            height 40px
            color red
            overflow hidden
            font-size 18px
        .list-item-bottom
          margin 10px 10px
          display flex
          flex-direction row
          justify-content space-around
          align-items center
          font-size 13px
          color #9b9898
          position relative
          .item-project
            position absolute
            left 10px
          .item-money
            position: absolute
            right 22px
            text-align center
            justify-content: flex-end
    .expense-tip
      text-align center
      color #9b9898
      line-height 10px
      font-size 12px
</style>
