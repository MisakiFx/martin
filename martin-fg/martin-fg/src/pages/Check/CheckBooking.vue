<template>
  <div class="check-booking">
    <div class="check-head">
      体检预约
    </div>
    <mt-checklist
      title="选择体检项目"
      v-model="projectValues"
      :options="checkProjectOptions"
      class="check-project">
    </mt-checklist>
    <mt-radio
      title="请选择体检日期(仅支持未来三天内预约)"
      v-model="checkDateValue"
      :options="checkDateOptions">
    </mt-radio>
    <mt-radio
      title="请选择体检时间"
      v-model="checkTimeValue"
      :options="checkTimeOptions">
    </mt-radio>
    <mt-radio
      title="请选择付款方式"
      v-model="payTypeValue"
      :options="payTypeOptions">
    </mt-radio>
    <button @click="postCheckBooking">预约</button>
  </div>
</template>

<script>
    import moment from 'moment'
    import {postCheckBooking} from '../../api'
    import {Indicator, MessageBox, Toast} from 'mint-ui'
    export default {
      name: "CheckBooking",
      mounted() {
        let now = moment()
        this.checkDateOptions = [
          {
            label: now.add(1,'d').format('Y年M月D日'),
            value: now.format('Y-MM-DD')
          },
          {
            label: now.add(1,'d').format('Y年M月D日'),
            value: now.format('Y-MM-DD')
          },
          {
            label: now.add(1,'d').format('Y年M月D日'),
            value: now.format('Y-MM-DD')
          }
        ]
      },
      methods:{
        async postCheckBooking() {
          //参数校验
          let projectValuesReal = []
          let money = 0
          for (let i in this.projectValues) {
            projectValuesReal.push(+this.projectValues[i])
            switch (+this.projectValues[i]) {
              case 1:
                money += 30
                break
              case 2:
                money += 30
                break
              case 3:
                money += 30
                break
              case 4:
                money += 199
                break
              case 5:
                money += 30
                break
              case 6:
                money += 30
                break
              case 7:
                money += 269
                break
              default :
                money += 0
            }
          }
          if (projectValuesReal.length === 0) {
            Toast("体检项目不能为空")
            return
          }
          if (!this.checkDateValue || !this.checkTimeValue) {
            Toast("请正确选择体检时间")
            return
          }
          if (!this.payTypeValue) {
            Toast("请选择付款方式")
            return
          }
          await this.$store.dispatch('reqUserExamination', this.$store.state.userInfo.open_id)
          money = money * this.$store.state.userExamination.card_type / 10
          MessageBox.confirm((+this.payTypeValue === 1 ? '预约将花费'+money+'元' : '预约将花费一次体检套餐') + ',是否确定预约?').then(async action=>{
            Indicator.open()
            let result = await postCheckBooking(this.$store.state.userInfo.open_id, {
              "check_project":projectValuesReal,
              "start_time":this.checkDateValue + ' ' + this.checkTimeValue,
              "pay_type":+this.payTypeValue
            })
            Indicator.close()
            if (result.code !== 0) {
              Toast(result.msg)
              return
            }
            Toast('预约成功')
            setTimeout(()=>{
              this.$router.replace('/me')
            }, 1000)
          })
        }
      },
      data() {
        return {
          checkTimeOptions:[
            {
              label: '8:00',
              value: '08:00:00'
            },
            {
              label: '10:00',
              value: '10:00:00'
            },
            {
              label: '14:00',
              value: '14:00:00'
            },
            {
              label: '16:00',
              value: '16:00:00'
            },
          ],
          checkProjectOptions:[
            {
              label: '内科(¥30)',
              value: '1'
            },
            {
              label: '外科(¥30)',
              value: '2'
            },
            {
              label: '耳鼻喉科(¥30)',
              value: '3'
            },
            {
              label: '肝功(¥199)',
              value: '4'
            },
            {
              label: '血糖(¥30)',
              value: '5'
            },
            {
              label: '血脂(¥30)',
              value: '6'
            },
            {
              label: '肾功($269)',
              value: '7'
            },
          ],
          checkDateOptions:[],
          payTypeOptions : [
            {
              label:'余额付款',
              value:'1',
            },
            {
              label:'体检套餐付款',
              value:'2',
            },
          ],
          projectValues : [],
          checkDateValue : '',
          checkTimeValue : '',
          payTypeValue : '',
        }
      }
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .check-booking
    width  100%
    height 100%
    .check-head
      width 100%
      height 50px
      line-height 50px
      font-weight bold
      text-align center
    button
      width 90%
      height 40px
      line-height 40px
      background-color: #e9232c
      text-align center
      margin 20px 5%
      border none
      font-size 16px
      color #fff
      border-radius 10px
</style>
