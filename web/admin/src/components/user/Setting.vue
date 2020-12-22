<template>
  <a-card>
    <h3>个人信息设置</h3>
    <a-form-model :model="profile" :hideRequiredMark="true">
      <a-form-model-item label="作者名">
        <a-input style="width:300px" v-model="profile.name"></a-input>
      </a-form-model-item>

      <a-form-model-item label="个人简介">
        <a-input type="textarea" v-model="profile.desc"></a-input>
      </a-form-model-item>

      <a-form-model-item label="QQ号码">
        <a-input style="width:300px" v-model="profile.qq_chat"></a-input>
      </a-form-model-item>

      <a-form-model-item label="微信号码">
        <a-input style="width:300px" v-model="profile.wechat"></a-input>
      </a-form-model-item>

      <a-form-model-item label="Email地址">
        <a-input style="width:300px" v-model="profile.email"></a-input>
      </a-form-model-item>

      <a-form-model-item label="微博账号">
        <a-input style="width:300px" v-model="profile.weibo"></a-input>
      </a-form-model-item>

      <a-form-model-item label="Bilibili主页">
        <a-input style="width:300px" v-model="profile.bili"></a-input>
      </a-form-model-item>

      <a-form-model-item label="个人介绍背景">
        <a-upload
          listType="picture"
          name="file"
          :action="upUrl"
          :headers="headers"
          @change="upChange"
        >
          <a-button>
            <a-icon type="upload" />点击上传
          </a-button>
          <br />
          <template v-if="profile.id">
            <img :src="profile.img" style="width:120px;height:100px" />
          </template>
        </a-upload>
      </a-form-model-item>

      <a-form-model-item label="个人照片">
        <a-upload
          listType="picture"
          name="file"
          :action="upUrl"
          :headers="headers"
          @change="upChange"
        >
          <a-button>
            <a-icon type="upload" />点击上传
          </a-button>
          <br />
          <template v-if="profile.id">
            <img :src="profile.avatar" style="width:120px;height:100px" />
          </template>
        </a-upload>
      </a-form-model-item>

      <a-form-model-item>
        <a-button type="danger" style="margin-right:15px" @click="Update">提交</a-button>
        <a-button type="primary" @click.once="$router.push('/index')">取消</a-button>
      </a-form-model-item>
    </a-form-model>
  </a-card>
</template>
<script>
import { Url } from '../../plugin/http'
export default {
  data() {
    return {
      profile: {
        id: 1,
        img: '',
        avatar: '',
        name: '',
        desc: '',
        qq_chat: '',
        email: '',
        weibo: '',
        bili: '',
        wechat: '',
      },
      upUrl: Url + 'upload',
      headers: '',
    }
  },
  created() {
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
    this.Getprofile()
  },
  methods: {
    // 获取个人设置信息
    async Getprofile() {
      const { data: res } = await this.$http.get('myprofile/1')
      this.profile = res.data
    },

    // 更新个人信息设置
    async Update() {
      const { data: res } = await this.$http.put('profile', this.profile)
      if (res.status !== 200) return this.$message.error(res.message)
      this.$message.success('更新信息成功')
      this.Getprofile()
    },
    // 上传图片
    upChange(info) {
      if (info.file.status !== 'uploading') {
      }
      if (info.file.status === 'done') {
        this.$message.success(`图片上传成功`)
        const imgUrl = info.file.response.url
        this.artInfo.img = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error(`图片上传失败`)
      }
    },
  },
}
</script>
<style lang="">
</style>