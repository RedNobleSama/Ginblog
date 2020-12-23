<template>
  <v-col>
    <v-card class="ma-3" v-for="item in artList" :key="item.id" link>
      <v-card-title>
        <v-chip color="indigo" class="mr-2 white--text font-weight-black">{{item.Category.name}}</v-chip>
        {{item.title}}
      </v-card-title>
      <v-card-text v-text="item.desc"></v-card-text>
      <v-card-text>
        <v-icon color="indigo darken-3" small>{{'mdi-clock-outline'}}</v-icon>
        {{item.CreatedAt}}
      </v-card-text>
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
      console.log(res)
    }
  }
}
</script>
<style lang="">
</style>