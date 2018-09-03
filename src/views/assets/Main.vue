<template>
<div class="animated fadeIn">

<b-container class="bv-example-row">
    <b-row>
        <b-col>
          <b-card v-for="item in investItems" :key="item.name" v-bind:title="item.name"
          v-bind:img-src="item.image"
          img-alt="Image"
          img-top
          tag="article"
          style="max-width: 20rem;"
          class="mb-2">
        <p class="card-text">
          {{item.description}}}
        </p>


        <router-link :to="{name: 'detail'}">
          <b-button variant="primary">Go</b-button>
        </router-link>



        </b-card>
        </b-col>

    <b-col>
          <b-card v-for="item in investItems1" :key="item.name" v-bind:title="item.name"
          v-bind:img-src="item.image"
          img-alt="Image"
          img-top
          tag="article"
          style="max-width: 20rem;"
          class="mb-2">
    <p class="card-text">
      {{item.description}}
    </p>
    <b-button href="#" variant="primary">Go</b-button>
  </b-card>
  </b-col>

  <b-col>
    <b-card v-for="item in investItems2" :key="item.name" v-bind:title="item.name"
          v-bind:img-src="item.image"
          img-alt="Image"
          img-top
          tag="article"
          style="max-width: 20rem;"
          class="mb-2">
    <p class="card-text">
      {{item.description}}}
    </p>
    <b-button href="#" variant="primary">Go</b-button>
  </b-card>
  </b-col>
    
    </b-row>
</b-container>


</div>
</template>

<script>
 import axios from 'axios'

export default {
  name: 'Main',

   mounted() {
    axios
      .get('http://localhost:3000/assets')
      .then(response => {
        this.investItems = response.data;
        this.investItems1 = response.data;
        this.investItems2 = response.data;
        });
        },
  data(){ return {
    investItems: [],
     investItems1: [],
    investItems2: []
  }
  }
}
</script>
