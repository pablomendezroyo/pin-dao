export const loopReload = () => {
  reloadPage();
};

const reloadPage = () => {
  setTimeout(() => {
    window.location.reload();
    reloadPage();
  }, 60000);
};
