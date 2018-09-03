<template>
<div class="animated fadeIn">

  <b-row>
      <b-col md="3" v-for="item in investItems" :key="item.name">
          <b-card  v-bind:title="item.name"
          v-bind:img-src="item.image"
          img-alt="Image"
          img-top
          tag="article"
          >
        <p class="card-text" style="height: 60px;">
          {{item.description}}}
        </p>


        <router-link :to="{name: 'detail'}">
          <b-button variant="primary">View <i class="far fa-eye"></i></b-button>
        </router-link>
        </b-card>
      </b-col>
  </b-row>


  <b-row>
      <b-col md="3" v-for="item in investItems1" :key="item.name">
          <b-card  v-bind:title="item.name"
          v-bind:img-src="item.image"
          img-alt="Image"
          img-top
          tag="article"
          >
        <p class="card-text" style="height: 60px;">
          {{item.description}}}
        </p>


        <router-link :to="{name: 'detail'}">
          <b-button variant="primary">View <i class="far fa-eye"></i></b-button> 
        </router-link>
        </b-card>
      </b-col>
  </b-row>
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
