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
              <td><span class="h5">25 USD</span></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <b-row>
      <b-col sm="12">
        <b-card>
          <div slot="header">
            <strong>Buy </strong>
          </div>

          <p> <strong>From:</strong> 21/12/2018  to 30/12/2018  </p>

          <p> <strong>Next round starts in:</strong> </p>

          <div>
            <b-progress :value="30" variant="info" striped :animated="animate" class="mb-2"></b-progress>
          </div>

        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import ColorTheme from './ColorTheme'

export default {
  name: 'round2',
}
</script>
