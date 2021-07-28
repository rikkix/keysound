<template>
  <div>
    <div class="flex flex-wrap p-4">
      <progress-bar :percentage="contentProgress" class="mx-2 mb-2 h-5" color="gray">
        <span class="text-xs text-white w-full flex justify-end pr-2">{{contentProgress}}%</span>
      </progress-bar>
    </div>
    <button @click="goQuestion"
            class="rounded-md border-2 border-gray-300 py-3 w-11/12 bg-gray-500 text-white font-bold" id="new"
            name="new" type="submit" v-if="completed"
            value="true">START
    </button>
    <FlashMessage :position="'right bottom'"/>
  </div>
</template>

<script>
  import ProgressBar from "../../components/ProgressBar";
  import axios from "axios";

  const qs = require('qs');

  export default {
    name: "Load",
    components: {
      ProgressBar
    },
    data() {
      return {
        contentProgress: 0,
        completed: false
      }
    },
    props: ['quizInfo'],
    methods: {
      goQuestion() {
        this.$router.push({
          name: 'Question',
          params: {
            quizInfo: this.quizInfo
          }
        })
      },
      processQuiz(data) {
        this.quizInfo = data;
        var base = data.base;
        var each = 100 / data.pieces.length;
        var itemsProcessed = 0;
        data.pieces.forEach((item) => {
          axios.get(base + "/" + item.path, {
            responseType: 'arraybuffer',
            headers: {
              'Content-Type': 'audio/mp3'
            }
          }).then((data) => {
            // console.log(data);
            var blob = new Blob([data.data], {
              type: 'audio/mp3'
            });
            sessionStorage.setItem(item.path, URL.createObjectURL(blob))
          }).then(() => {
            this.contentProgress += each;
            itemsProcessed++;
            if (itemsProcessed === data.pieces.length) {
              this.completed = true
            }
          }).catch((error) => {
            console.log(error);
            this.flashMessage.error({
              title: error.name + " (" + error.config.url + ")",
              message: error.message,
            });
          })
        })
      }
    },
    mounted() {
      if (this.quizInfo == null) {
        axios.post("/api/get", qs.stringify({'id': this.$route.params.id})).then((res) => {
          return res.data
        }).then((data) => this.processQuiz(data))
          .catch((error) => {
            console.log(error);
            this.flashMessage.error({
              title: error.name + " (" + error.config.url + ")",
              message: error.message,
            });
          })
      } else {
        this.processQuiz(this.quizInfo)
      }
    }
  }
</script>

<style scoped>

</style>

<style>
  /*.bg-dark-500*/
</style>
