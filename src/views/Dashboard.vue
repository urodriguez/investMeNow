<template>
  <div class="animated fadeIn">
  <b-row>
      <b-col sm="12" lg="12">
        <b-card>
        <b-row>
          <b-col>
            <h4 id="traffic" class="card-title mb-0">Variants of Purchases and Sales</h4>
          </b-col>
        </b-row>
        <main-chart-example chartId="main-chart-01" class="chart-wrapper" style="height:300px;margin-top:40px;" height="300"></main-chart-example>
        <div slot="footer">
          <b-row class="text-center">
            <b-col class="mb-sm-2 mb-0">
              <div class="text-muted">Purchases</div>
              <strong>29.703 (18%)</strong>
              <b-progress height={} class="progress-xs mt-2" :precision="1" variant="success" :value="40"></b-progress>
            </b-col>
            <b-col class="mb-sm-2 mb-0">
              <div class="text-muted">Sales</div>
              <strong>22.123 (63%)</strong>
              <b-progress height={} class="progress-xs mt-2" :precision="1" variant="danger" :value="80"></b-progress>
            </b-col>
          </b-row>
        </div>
      </b-card>
    </b-col>
  </b-row>

  <b-row>
    <b-col sm="6" lg="6">
      <b-card title="Buy" img-alt="Image" img-top tag="article">
        <div class="divTable blueTable">
          <div class="divTableBody">
            <div class="divTableRow">
              <div class="divTableCell">15.09016401</div>
              <div class="divTableCell">13.98245161</div>
              <div class="divTableCell">12:05h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">11.98245161</div>
              <div class="divTableCell">16.00098766</div>
              <div class="divTableCell">12:10h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">21.67856734</div>
              <div class="divTableCell">14.12312223</div>
              <div class="divTableCell">12:15h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">15.09016401</div>
              <div class="divTableCell">13.98245161</div>
              <div class="divTableCell">12:05h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">11.98245161</div>
              <div class="divTableCell">16.00098766</div>
              <div class="divTableCell">12:10h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">21.67856734</div>
              <div class="divTableCell">14.12312223</div>
              <div class="divTableCell">12:15h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">15.09016401</div>
              <div class="divTableCell">13.98245161</div>
              <div class="divTableCell">12:05h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">11.98245161</div>
              <div class="divTableCell">16.00098766</div>
              <div class="divTableCell">12:10h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">21.67856734</div>
              <div class="divTableCell">14.12312223</div>
              <div class="divTableCell">12:15h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">15.09016401</div>
              <div class="divTableCell">13.98245161</div>
              <div class="divTableCell">12:05h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">11.98245161</div>
              <div class="divTableCell">16.00098766</div>
              <div class="divTableCell">12:10h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">21.67856734</div>
              <div class="divTableCell">14.12312223</div>
              <div class="divTableCell">12:15h 03/09/2018</div>
            </div>
          </div>
        </div>
      </b-card>
    </b-col>
    <b-col sm="6" lg="6">
      <b-card title="Sale" img-alt="Image" img-top tag="article">
        <div class="divTable blueTable">
          <div class="divTableBody">
            <div class="divTableRow">
              <div class="divTableCell">15.09016401</div>
              <div class="divTableCell">13.98245161</div>
              <div class="divTableCell">12:05h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">11.98245161</div>
              <div class="divTableCell">16.00098766</div>
              <div class="divTableCell">12:10h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">21.67856734</div>
              <div class="divTableCell">14.12312223</div>
              <div class="divTableCell">12:15h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">15.09016401</div>
              <div class="divTableCell">13.98245161</div>
              <div class="divTableCell">12:05h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">11.98245161</div>
              <div class="divTableCell">16.00098766</div>
              <div class="divTableCell">12:10h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">21.67856734</div>
              <div class="divTableCell">14.12312223</div>
              <div class="divTableCell">12:15h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">15.09016401</div>
              <div class="divTableCell">13.98245161</div>
              <div class="divTableCell">12:05h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">11.98245161</div>
              <div class="divTableCell">16.00098766</div>
              <div class="divTableCell">12:10h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">21.67856734</div>
              <div class="divTableCell">14.12312223</div>
              <div class="divTableCell">12:15h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">15.09016401</div>
              <div class="divTableCell">13.98245161</div>
              <div class="divTableCell">12:05h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">11.98245161</div>
              <div class="divTableCell">16.00098766</div>
              <div class="divTableCell">12:10h 03/09/2018</div>
            </div>
            <div class="divTableRow">
              <div class="divTableCell">21.67856734</div>
              <div class="divTableCell">14.12312223</div>
              <div class="divTableCell">12:15h 03/09/2018</div>
            </div>
          </div>
        </div>
      </b-card>
    </b-col>
  </b-row>

  <b-row>
    <b-col sm="6" lg="6">
      <b-card>
        <b-row>
    <b-col sm="6" lg="6">
      
      
      <b-form @submit="onSubmit" @reset="onReset" v-if="show">
        <b-form-group id="exampleInputGroup1"
                      label="Price"
                      label-for="exampleInput1">
          <b-form-input id="exampleInput1"
                        type="text"
                        required
                        placeholder="Price">
          </b-form-input>
        </b-form-group>
        <b-form-group id="exampleInputGroup2"
                      label="Amount"
                        label-for="exampleInput2">
          <b-form-input id="exampleInput2"
                        type="text"
                        required
                        placeholder="Amount">
          </b-form-input>
        </b-form-group>
        <b-button type="submit" variant="primary">Buy <i class="fas fa-angle-double-up"></i></b-button>
      </b-form>
    </b-col>
    <b-col sm="6" lg="6">
      <b-form @submit="onSubmit" @reset="onReset" v-if="show">
        <b-form-group id="exampleInputGroup1"
                      label="Price"
                      label-for="exampleInput1">
          <b-form-input id="exampleInput1"
                        type="text"
                        required
                        placeholder="Price">
          </b-form-input>
        </b-form-group>
        <b-form-group id="exampleInputGroup2"
                      label="Amount"
                        label-for="exampleInput2">
          <b-form-input id="exampleInput2"
                        type="text"
                        required
                        placeholder="Amount">
          </b-form-input>
        </b-form-group>
        <b-button type="submit" variant="danger">Sell <i class="fas fa-angle-double-down"></i></b-button>
      </b-form>
    </b-col>
  </b-row>
      </b-card>
    </b-col>
  </b-row>

  
  </div>
</template>

<script>
import CardLine1ChartExample from './dashboard/CardLine1ChartExample'
import CardLine2ChartExample from './dashboard/CardLine2ChartExample'
import CardLine3ChartExample from './dashboard/CardLine3ChartExample'
import CardBarChartExample from './dashboard/CardBarChartExample'
import MainChartExample from './dashboard/MainChartExample'
import SocialBoxChartExample from './dashboard/SocialBoxChartExample'
import CalloutChartExample from './dashboard/CalloutChartExample'
import { Callout } from '@coreui/vue'

export default {
  name: 'dashboard',
  components: {
    Callout,
    CardLine1ChartExample,
    CardLine2ChartExample,
    CardLine3ChartExample,
    CardBarChartExample,
    MainChartExample,
    SocialBoxChartExample,
    CalloutChartExample
  },
  data: function () {
    return {
      selected: 'Month',
      tableItems: [
        {
          avatar: { url: 'img/avatars/1.jpg', status: 'success' },
          user: { name: 'Yiorgos Avraamu', new: true, registered: 'Jan 1, 2015' },
          country: { name: 'USA', flag: 'us' },
          usage: { value: 50, period: 'Jun 11, 2015 - Jul 10, 2015' },
          payment: { name: 'Mastercard', icon: 'fa fa-cc-mastercard' },
          activity: '10 sec ago'
        },
        {
          avatar: { url: 'img/avatars/2.jpg', status: 'danger' },
          user: { name: 'Avram Tarasios', new: false, registered: 'Jan 1, 2015' },
          country: { name: 'Brazil', flag: 'br' },
          usage: { value: 22, period: 'Jun 11, 2015 - Jul 10, 2015' },
          payment: { name: 'Visa', icon: 'fa fa-cc-visa' },
          activity: '5 minutes ago'
        },
        {
          avatar: { url: 'img/avatars/3.jpg', status: 'warning' },
          user: { name: 'Quintin Ed', new: true, registered: 'Jan 1, 2015' },
          country: { name: 'India', flag: 'in' },
          usage: { value: 74, period: 'Jun 11, 2015 - Jul 10, 2015' },
          payment: { name: 'Stripe', icon: 'fa fa-cc-stripe' },
          activity: '1 hour ago'
        },
        {
          avatar: { url: 'img/avatars/4.jpg', status: '' },
          user: { name: 'Enéas Kwadwo', new: true, registered: 'Jan 1, 2015' },
          country: { name: 'France', flag: 'fr' },
          usage: { value: 98, period: 'Jun 11, 2015 - Jul 10, 2015' },
          payment: { name: 'PayPal', icon: 'fa fa-paypal' },
          activity: 'Last month'
        },
        {
          avatar: { url: 'img/avatars/5.jpg', status: 'success' },
          user: { name: 'Agapetus Tadeáš', new: true, registered: 'Jan 1, 2015' },
          country: { name: 'Spain', flag: 'es' },
          usage: { value: 22, period: 'Jun 11, 2015 - Jul 10, 2015' },
          payment: { name: 'Google Wallet', icon: 'fa fa-google-wallet' },
          activity: 'Last week'
        },
        {
          avatar: { url: 'img/avatars/6.jpg', status: 'danger' },
          user: { name: 'Friderik Dávid', new: true, registered: 'Jan 1, 2015' },
          country: { name: 'Poland', flag: 'pl' },
          usage: { value: 43, period: 'Jun 11, 2015 - Jul 10, 2015' },
          payment: { name: 'Amex', icon: 'fa fa-cc-amex' },
          activity: 'Last week'
        }
      ],
      tableFields: {
        avatar: {
          label: '<i class="icon-people"></i>',
          class: 'text-center'
        },
        user: {
          label: 'User'
        },
        country: {
          label: 'Country',
          class: 'text-center'
        },
        usage: {
          label: 'Usage'
        },
        payment: {
          label: 'Payment method',
          class: 'text-center'
        },
        activity: {
          label: 'Activity'
        }
      },
      form: {
        email: '',
        name: '',
        food: null,
        checked: []
      },
      foods: [
        { text: 'Select One', value: null },
        'Carrots', 'Beans', 'Tomatoes', 'Corn'
      ],
      show: true
    }
  },
  methods: {
    variant (value) {
      let $variant
      if (value <= 25) {
        $variant = 'info'
      } else if (value > 25 && value <= 50) {
        $variant = 'success'
      } else if (value > 50 && value <= 75) {
        $variant = 'warning'
      } else if (value > 75 && value <= 100) {
        $variant = 'danger'
      }
      return $variant
    },
    flag (value) {
      return 'flag-icon flag-icon-' + value
    },
    onSubmit (evt) {
      evt.preventDefault();
      alert(JSON.stringify(this.form));
    },
    onReset (evt) {
      evt.preventDefault();
      /* Reset our form values */
      this.form.email = '';
      this.form.name = '';
      this.form.food = null;
      this.form.checked = [];
      /* Trick to reset/clear native browser form validation state */
      this.show = false;
      this.$nextTick(() => { this.show = true });
    }
  }
}
</script>

<style>
  /* IE fix */
  #card-chart-01, #card-chart-02 {
    width: 100% !important;
  }
</style>
