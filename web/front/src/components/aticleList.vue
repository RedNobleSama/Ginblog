<template>
  <v-col>
    <v-card
      class="ma-3"
      v-for="item in artList"
      :key="item.id"
      link
      @click="$router.push(`/details/${item.ID}`)"
    >
      <v-row no-gutters>
        <v-col class="d-flex justify-center align-center mx-3" cols="1">
          <v-img :src="item.img" max-height="100" max-width="100"></v-img>
        </v-col>
        <v-col>
          <v-card-title>
            <v-chip
              color="pink"
              label
              class="mr-2 white--text font-weight-black"
            >{{item.Category.name}}</v-chip>
            {{item.title}}
          </v-card-title>
          <v-card-subtitle v-text="item.desc"></v-card-subtitle>
          <v-divider></v-divider>
          <v-card-text>
            <v-icon color="red darken-2" small>{{'mdi-calendar-month-outline'}}</v-icon>
            <span class="subheading ma-1 py-1">{{item.CreatedAt}}</span>
          </v-card-text>
        </v-col>
      </v-row>
    </v-card>
    <div class="text-center">
      <v-pagination
        v-model="queryParam.pagenum"
        :length="Math.ceil(this.total / this.queryParam.pagesize)"
        :total-visible="7"
        @input="getArtList()"
      ></v-pagination>
    </div>
  </v-col>
</template>
<script>
export default {
  data() {
    return {
      artList: [],
      total: 0,
      queryParam: {
        pagesize: 7,
        pagenum: 1
      }
    }
  },
  created() {
    this.getArtList()
  },
  filters: {},

  methods: {
    async getArtList() {
      const { data: res } = await this.$http.get('article', {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      this.artList = res.data
      this.total = res.total
    }
  }
}
</script>
<style lang="">
</style>