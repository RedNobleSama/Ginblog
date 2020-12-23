<template>
  <v-card class="mx-auto" max-width="350">
    <v-img src="../assets/bluebg.png">
      <v-card-title>
        <v-col align="center">
          <v-avatar size="130" color="grey">
            <img src="../assets/123.jpg" alt="John" />
          </v-avatar>
          <div class="ma-3 white--text">{{profile.name}}</div>
        </v-col>
      </v-card-title>

      <v-divider color="white"></v-divider>
    </v-img>

    <v-card-title>About Me:</v-card-title>
    <v-card-text>{{profile.desc}}</v-card-text>
    <v-divider color="indigo"></v-divider>

    <v-list dense>
      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="primary">{{'mdi-qqchat'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profile.qq_chat}}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="green darken-2">{{'mdi-wechat'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profile.wechat}}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="deep-orange">{{'mdi-sina-weibo'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profile.weibo}}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="blue darken-3">{{'mdi-youtube-tv'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profile.bili}}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="yellow darken-3">{{'mdi-email'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profile.email}}</v-list-item-content>
      </v-list-item>
    </v-list>
  </v-card>
</template>
<script>
export default {
  data: () => ({
    profile: {}
  }),
  created() {
    this.Getprofile()
  },
  methods: {
    async Getprofile() {
      const { data: res } = await this.$http.get('myprofile/1')
      this.profile = res.data
    }
  }
}
</script>
<style lang="">
</style>