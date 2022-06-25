<template>
  <div>
    <el-calendar v-model="date" />
  </div>
</template>

<script>
export default {
  name:"IndexItem",
  components: {},
  data() {
    return {
      date:new Date()
    }
  },
  methods: {},
  mounted () {},
}

</script>
<style lang='scss' scoped>

</style>