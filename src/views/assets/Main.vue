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
export default {
  name: 'Main',
  data(){ return {
    investItems: [
      { name: 'Real State / Buildings', description: 'Real State / Buildings Real State / Buildings Real State / Buildings Real State / Buildings', image:"https://images.pexels.com/photos/290275/pexels-photo-290275.jpeg?cs=srgb&dl=city-skyline-buildings-290275.jpg&fm=jpg" },
      { name: 'Copyrights', description: 'Copyrights Copyrights Copyrights Copyrights', image:'https://sites.google.com/site/derechosdedeautorr/_/rsrc/1468755296557/para-que-sirve/11504987-vector-sello-danado-los-derechos-de-autor.jpg' },
      { name: 'Independent Project', description: 'Independent Project Independent Project Independent Project Independent Project', image:'http://factorypyme.thestandardit.com/wp-content/uploads/sites/5/2014/09/Adquisicion.jpg' },
      { name: 'Next Sport Stars', description: 'Next Sport Stars Next Sport Stars Next Sport Stars Next Sport Stars', image:'https://pbs.twimg.com/profile_images/875393235354439680/qWkpDQnh_400x400.jpg' },
    ],
     investItems1: [
      { name: 'Real State / Buildings', description: 'Real State / Buildings Real State / Buildings Real State / Buildings Real State / Buildings', image:"https://images.pexels.com/photos/290275/pexels-photo-290275.jpeg?cs=srgb&dl=city-skyline-buildings-290275.jpg&fm=jpg" },
      { name: 'Copyrights', description: 'Copyrights Copyrights Copyrights Copyrights', image:'https://sites.google.com/site/derechosdedeautorr/_/rsrc/1468755296557/para-que-sirve/11504987-vector-sello-danado-los-derechos-de-autor.jpg' },
      { name: 'Independent Project', description: 'Independent Project Independent Project Independent Project Independent Project', image:'http://factorypyme.thestandardit.com/wp-content/uploads/sites/5/2014/09/Adquisicion.jpg' },
      { name: 'Next Sport Stars', description: 'Next Sport Stars Next Sport Stars Next Sport Stars Next Sport Stars', image:'https://pbs.twimg.com/profile_images/875393235354439680/qWkpDQnh_400x400.jpg' },
    ],
    investItems2: [
      { name: 'Real State / Buildings', description: 'Real State / Buildings Real State / Buildings Real State / Buildings Real State / Buildings', image:"https://images.pexels.com/photos/290275/pexels-photo-290275.jpeg?cs=srgb&dl=city-skyline-buildings-290275.jpg&fm=jpg" },
      { name: 'Copyrights', description: 'Copyrights Copyrights Copyrights Copyrights', image:'https://sites.google.com/site/derechosdedeautorr/_/rsrc/1468755296557/para-que-sirve/11504987-vector-sello-danado-los-derechos-de-autor.jpg' },
      { name: 'Independent Project', description: 'Independent Project Independent Project Independent Project Independent Project', image:'http://factorypyme.thestandardit.com/wp-content/uploads/sites/5/2014/09/Adquisicion.jpg' },
      { name: 'Next Sport Stars', description: 'Next Sport Stars Next Sport Stars Next Sport Stars Next Sport Stars', image:'https://pbs.twimg.com/profile_images/875393235354439680/qWkpDQnh_400x400.jpg' },
    ]
  }
  }
}
</script>
