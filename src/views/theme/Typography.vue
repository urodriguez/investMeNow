<template>
  <div class="animated fadeIn">
    <div class="card">
      <div class="card-header">
        Round Information
      </div>

      <div class="card-body">
        <p>In the next table you can check the current price for one InvestMoney (IM)</p>
        <p>(max amount for this round is: <strong>1000 IM</strong>)</p>
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
              <td><span class="h5">100 USD</span></td>
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
          <b-row>
            <b-col sm="2">
              <b-form-group>
                <b-form-input type="text" id="name" placeholder="Enter your IM amount"></b-form-input>
              </b-form-group>
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
  </div>
</template>

<script>
export default {
  name: 'typography'
}
</script>
