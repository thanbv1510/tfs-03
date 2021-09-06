<template>
  <div class="modal">
    <div class="modal-content">
      <button class="close" @click="closeModal" type="button">X</button>
      <p>Success</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ModalItem',
  methods: {
    closeModal() {
      this.$emit('closeModalEvt')
    }
  }
}
</script>

<style scoped>
.modal {
  position: relative;
  background: rgba(0, 0, 0, 0.6);
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

.modal-content {
  text-align: center;
  width: 15%;
  height: 10%;
  margin: 20% auto;
  background-color: rgb(255, 255, 255);
  border-radius: 15px;
  font-size: 4rem;
  color: darkgreen;
  position: relative;

  display: flex;
  justify-content: center;
  align-items: center;
}

.modal .close {
  position: absolute;
  top: 5%;
  right: 5%;
  width: 5%;
  border-radius: 5px;
  font-size: 2rem;
  background-color: rgb(255, 255, 255);
  font-weight: bold;
  cursor: pointer;
}

.modal .close:hover {
  color: red;
}
</style>