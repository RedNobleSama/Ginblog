<template>
  <v-card class="mx-auto" max-width="300">
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

    <v-list nav dense>
      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon>{{'mdi-qqchat'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profile.qq_chat}}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon>{{'mdi-wechat'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profile.wechat}}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon>{{'mdi-sina-weibo'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">12334566</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon>{{'mdi-youtube-tv'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profile.bili}}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon>{{'mdi-twitter'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">12334566</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon>{{'mdi-email'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">12334566</v-list-item-content>
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