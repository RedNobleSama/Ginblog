<template>
  <div>
    <div class="d-flex pa-3 justify-center text-h3 font-weight-bold">{{artInfo.title}}</div>
    <v-divider></v-divider>
    <v-chip
      class="mt-4 ml-3 pa-2 font-weight-medium"
      color="blue"
      text-color="white"
      label
    >{{artInfo.desc}}</v-chip>
    <p class="ma-5 pa-3 text-justify" v-html="artInfo.content"></p>
  </div>
</template>
<script>
export default {
  props: ['id'],
  data() {
    return {
      artInfo: {}
    }
  },
  created() {
    this.getArtInfo()
  },
  methods: {
    // 查询文章信息
    async getArtInfo() {
      const { data: res } = await this.$http.get(`article/info/${this.id}`)
      this.artInfo = res.data
      console.log(this.artInfo)
    }
  }
}
</script>
<style lang="">
</style>