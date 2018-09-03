<template>
  <div class="wrapper">
    <div class="animated fadeIn">
      <b-row>
        <b-col md="12" >
          <b-card
            header-tag="header"
            footer-tag="footer">
            <div slot="header">
              <strong>Building Images</strong>
              <div class="card-header-actions">
                <a href="https://bootstrap-vue.js.org/docs/components/carousel" class="card-header-action" rel="noreferrer noopener" target="_blank">
                </a>
              </div>
            </div>
            <div>
              <b-carousel id="carousel1"
                          style="text-shadow: 1px 1px 2px #333;"
                          controls
                          indicators
                          background="#ababab"
                          :interval="4000"
                          img-width="1024"
                          img-height="480"
                          v-model="slide"
                          @sliding-start="onSlideStart"
                          @sliding-end="onSlideEnd"
              >

                <!-- Text slides with image -->
                <b-carousel-slide img-src="https://roofstock-cdn.azureedge.net/public/properties/1675745/photo/b166a2f7-46cf-4042-81f4-9bd0847adc56_1._Front-modified-20181715154410_1280.jpg"
                ></b-carousel-slide>

                <!-- Slides with custom text -->
                <b-carousel-slide img-src="https://roofstock-cdn.azureedge.net/public/properties/1675745/photo/1aa317db-5007-4ccb-b5fb-b69390b5d1f3_2._Rear-modified-20181715154450_1280.jpg">
                </b-carousel-slide>

                <!-- Slides with image only -->
                <b-carousel-slide img-src="https://roofstock-cdn.azureedge.net/public/properties/1675745/photo/4cfec31f-a50b-4f22-a1e6-c4c2bf5046d4_3.1_Kitchen-modified-20181715155240_1280.jpg">
                </b-carousel-slide>

                <!-- Slides with image only -->
                <b-carousel-slide img-src="https://roofstock-cdn.azureedge.net/public/properties/1675745/photo/edd713f9-88d3-480f-9238-19c721a58fde_4_-_Floor_Plan-modified-20181717134057_1280.jpg">
                </b-carousel-slide>

                <!-- Slides with image only -->
                <b-carousel-slide img-src="https://roofstock-cdn.azureedge.net/public/properties/1675745/photo/b559d975-3cb4-474e-8aae-52d76b506cd2_5.1_Kitchen-modified-20181715154646_1280.jpg">
                </b-carousel-slide>

                <!-- Slides with image only -->
                <b-carousel-slide img-src="https://roofstock-cdn.azureedge.net/public/properties/1675745/photo/f67f6fc4-7eef-48ed-a38a-1df7bbe5b329_7._Bathroom-modified-20181715154838_1280.jpg">
                </b-carousel-slide>

              </b-carousel>
            </div>
          </b-card>
        </b-col>

      </b-row>

      <b-row>
        <b-col md="6" >

          <b-card header-tag="header" footer-tag="footer">
            <div slot="header">
              <strong>Invest Information</strong>
            </div>

              <h2>Initial Investment <i class="fas fa-question-circle" v-b-tooltip.hover title="I am a tooltip!"></i></h2> 
              <br>
              <h3>$52,565</h3>
              <br>
              <h4>List Price:  <strong> $192,900 </strong> </h4>
          </b-card>

        </b-col>

        <b-col md="6" >
          <b-card header-tag="header" footer-tag="footer">
            <div slot="header">
              <strong>Location</strong>
            </div>

            <div class="google-map">
              <iframe width="100%" frameborder="0" scrolling="no" marginheight="0" marginwidth="0" src="https://maps.google.co.uk/maps?f=q&source=s_q&hl=en&geocode=&q=15+Springfield+Way,+Hythe,+CT21+5SH&aq=t&sll=52.8382,-2.327815&sspn=8.047465,13.666992&ie=UTF8&hq=&hnear=15+Springfield+Way,+Hythe+CT21+5SH,+United+Kingdom&t=m&z=14&ll=51.077429,1.121722&output=embed"></iframe>
            </div>
          </b-card>
        </b-col>
      </b-row>

      <b-row>
        <b-col md="12" >
        <b-card>
          <div slot="header">
            <strong>Property Management Information</strong>
          </div>

          <b-row>

            <b-col sm="6">
              <p>Korn Buildings</p>
              <p>(123) 456-7890</p>
              <p>test@test.com</p>
              <p>4435 Sumner Ave</p>
              <p>Memphis, TN 38122</p>
            </b-col>

            <b-col sm="6">
              <img class="navbar-brand-full" src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSdS1Wxk6dk4cCFlOV2IRofHTL-ho19zdx6xwk0sT5javgC2p63" style="max-width:150px; max-height:150px;">
            </b-col>

          </b-row>

        </b-card>
        </b-col>
      </b-row>

      <b-row>
        <b-col md="12" >
        <b-card class="warning">
          <div slot="header">
            <strong>Tips for Self Managing Properties</strong>
          </div>

          <b-row>

            <b-col sm="4">
              <p>You’ve invested in a Rental Property, and are thinking about managing the property yourself? Here are the items you need to be aware of: Renting, Operating, Risk Management</p>
            </b-col>

            <b-col sm="4">
              <p>Renting</p>
              <ul>
                <li>
                  Get property ready
                </li>
                <li>
                  Advertise
                </li>
                <li>
                  Show the property
                </li>
                <li>
                  Applicant screening
                </li>
                <li>
                  Move in tenants
                </li>
              </ul>
            </b-col>


            <b-col sm="4">
              <p>Risk Management</p>
              <ul>
                <li>
                  Fair housing
                </li>
                <li>
                  Insurance liability
                </li>
                <li>
                  Property damage
                </li>
                <li>
                  Habitability issues
                </li>
                <li>
                  Legal costs
                </li>
              </ul>
            </b-col>
          </b-row>

        </b-card>
        </b-col>
      </b-row>


      <b-row>
        <b-col md="12" >
        <b-card class="warning">
          <div slot="header">
            <strong>Tips for Self Managing Properties</strong>
          </div>

          <b-row>

            <b-col sm="4">
              <p>You’ve invested in a Rental Property, and are thinking about managing the property yourself? Here are the items you need to be aware of: Renting, Operating, Risk Management</p>
            </b-col>

            <b-col sm="4">
              <p>Renting</p>
              <ul>
                <li>
                  Get property ready
                </li>
                <li>
                  Advertise
                </li>
                <li>
                  Show the property
                </li>
                <li>
                  Applicant screening
                </li>
                <li>
                  Move in tenants
                </li>
              </ul>
            </b-col>

          </b-row>

        </b-card>
        </b-col>
      </b-row>



      <b-button pressed block variant="dark" class="btn-pill" style="height: 64px" aria-pressed="true">
        <h1>Invest Me Now</h1>
      </b-button>
      <br>
      <br>
      <br>
      <br>




    </div>
  </div>
</template>

<script>
export default {
  name: 'carousels',
  data () {
    return {
      slide: 0,
      sliding: null
    }
  },
  methods: {
    onSlideStart (slide) {
      console.log('onSlideStart', slide)
      this.sliding = true
    },
    onSlideEnd (slide) {
      console.log('onSlideEnd', slide)
      this.sliding = false
    }
  }
}
</script>
