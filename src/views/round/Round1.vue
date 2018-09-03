<template>
  <div class="animated fadeIn">
    <div class="card">
      <div class="card-header">
        Round Information
      </div>

      <div class="card-body">
        <p>In the next table you can check the current price for one InvestMoney (IM)</p>
        <p>(max amount for this round is: <strong>500.000 IM</strong>)</p>
        <table class="table">
          <thead>
            <tr>
              <th>Amount</th>
              <th>Price</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>
                <p>1</p>
              </td>
              <td><span class="h5">10 USD</span></td>
            </tr>
          </tbody>
        </table>
        
        <h4> <strong>Total founded:</strong> </h4>

        <div>
          <b-progress height="32px" class="mt-1" :max="max" show-value>
            <b-progress-bar :value="counter" variant="success">
              <strong style="font-size: 16px">{{ counter }} / {{ max }}</strong>
            </b-progress-bar>
          </b-progress>
        </div>

      </div>
    </div>

    <b-row>
      <b-col sm="12">
        <b-card>
          <div slot="header">
            <strong>Buy </strong>
          </div>
          <b-row>
            <b-col sm="2">
              <b-form-group>
                <b-form-input type="text" id="name" placeholder="Enter your IM amount"></b-form-input>
              </b-form-group>
            </b-col>

            <b-col sm="12">
              <b-card>
                <div slot="header">
                  <strong>Credit Card </strong>
                </div>
                <b-row>
                  <b-col sm="12">
                    <b-form-group>
                      <label for="name">Name</label>
                      <b-form-input type="text" id="name" placeholder="Enter your name"></b-form-input>
                    </b-form-group>
                  </b-col>
                </b-row>
                <b-row>
                  <b-col sm="12">
                    <b-form-group>
                      <label for="ccnumber">Credit Card Number</label>
                      <b-form-input type="text" id="ccnumber" placeholder="0000 0000 0000 0000"></b-form-input>
                    </b-form-group>
                  </b-col>
                </b-row>
                <b-row>
                  <b-col sm="4">
                    <b-form-group >
                      <label for="month1">Month</label>
                      <b-form-select id="month1"
                        :plain="true"
                        :options="[1,2,3,4,5,6,7,8,9,10,11,12]">
                      </b-form-select>
                    </b-form-group>
                  </b-col>
                  <b-col sm="4">
                    <b-form-group>
                      <label for="year1">Year</label>
                      <b-form-select id="year1"
                        :plain="true"
                        :options="[2014,2015,2016,2017,2018,2019,2020,2021,2022,2023,2024,2025]">
                      </b-form-select>
                    </b-form-group>
                  </b-col>
                  <b-col sm="4">
                    <b-form-group>
                      <label for="cvv">CVV/CVC</label>
                      <b-form-input type="text" id="cvv" placeholder="123"></b-form-input>
                    </b-form-group>
                  </b-col>
                </b-row>
              </b-card>
            </b-col>

            <b-col sm="1">
              <b-form-group>
                <b-button pressed block variant="dark" class="btn-pill" aria-pressed="true">Buy</b-button>
              </b-form-group>
              
            </b-col>
          </b-row>
        </b-card>
      </b-col>
    </b-row>

    <b-row>
      
    </b-row>
  </div>
</template>

<script>
export default {
  name: 'round1',
  data () {
    return {
      counter: 8750000,
      max: 10000000,
      max2: 50,
      value: 33.333333333,
      value3: 75,
      bars: [
        {variant: 'success', value: 75},
        {variant: 'info', value: 75},
        {variant: 'warning', value: 75},
        {variant: 'danger', value: 75},
        {variant: 'primary', value: 75},
        {variant: 'secondary', value: 75},
        {variant: 'dark', value: 75}
      ],
      timer: null,
      striped: true,
      animate: true,
      max3: 100,
      values: [ 15, 30, 20 ]
    }
  },
  methods: {
    clicked () {
      this.counter = Math.random() * this.max
      console.log('Change progress to ' +
          Math.round(this.counter * 100) / 100)
    }
  },
  mounted () {
    this.timer = setInterval(() => {
      this.bars.forEach(bar => {
        bar.value = 25 + (Math.random() * 75)
      })
    }, 2000)
  },
  beforeDestroy () {
    clearInterval(this.timer)
    this.timer = null
  }
}
</script>
