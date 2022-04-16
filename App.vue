<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png" />
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <!-- <pre>{{ $v }}</pre> -->
    <!-- <form @submit.prevent="save()">
      <label>Name:</label>
      <input
        @blur="$v.formData.basicDetails.name.$touch()"
        type="text"
        v-model="formData.basicDetails.name"
      />
      <div
        v-if="
          (issubmit || $v.formData.basicDetails.name.$dirty) && !$v.formData.basicDetails.name.required
        "
      >
        It is required field
      </div>
      <label>Mobile No :</label>
      <input
        type="text"
        @blur="$v.formData.basicDetails.mobileNo.$touch()"
        v-model="formData.basicDetails.mobileNo"
      />
      <div
        v-if="
          (issubmit || $v.formData.basicDetails.mobileNo.$dirty) &&
          !$v.formData.basicDetails.mobileNo.required
        "
      >
        It is required field
      </div>
      <button type="submit">Click</button>

    </form> -->
    Files:
    <input
      type="file"
      @change="uploadFile"
      ref="file"
    /><br /><br />
    <button @click="submitFile">Upload!</button>
  </div>
</template>

<script>
import swal from "sweetalert";
import axios from "axios";

// import HelloWorld from './components/HelloWorld.vue'
import { required } from "vuelidate/lib/validators";

export default {
  name: "App",
  components: {
    // HelloWorld
  },

  data() {
    return {
      issubmit: false,
      image: null,
      formData: {
        basicDetails: {
          name: "",
          mobileNo: "",
        },
      },
    };
  },
  validations() {
    return {
      formData: {
        basicDetails: {
          mobileNo: {
            required,
          },
          name: {
            required,
          },
        },
      },
    };
  },
  methods: {
    uploadFile() {
      this.image = this.$refs.file.files[0];
      console.log("..", this.image);
    },
    submitFile(){
       const formData = new FormData();
        formData.append('file', this.image);
        axios.post("http://localhost:1323/upload",formData).then(res=>{
          console.log(res.data);
        })

    },
    save() {
      this.issubmit = true;

      if (this.issubmit && !this.$v.$invalid) {
        console.log("formdata", this.formData);
        swal({
          title: "Good job!",
          text: "You clicked the button!",
          icon: "success",
          button: "Aww yiss!",
        });
        this.formData = {};
      }
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
