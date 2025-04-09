document.addEventListener('DOMContentLoaded', function() {
    const categoriesContainer = document.querySelector('.categories-container');
    if (categoriesContainer) {
        categoriesContainer.addEventListener('wheel', function(e) {
            if (e.deltaY !== 0) {
                this.scrollTop += e.deltaY;
                e.preventDefault();
            }
        });
    }
});