function isVisible() {
    let isSwiped = false;
    const header = document.querySelector('.header');
    const contentBox = document.querySelector('.content_box');
    const description = document.querySelector('.about_us_description');
    const img = document.querySelector('.about_us_img');

    window.addEventListener('wheel', function (event) {
        if (event.deltaY > 0 && !isSwiped) {
            isSwiped = true;
            header.style.transform = 'translateY(-950px)';
            contentBox.style.transform = 'translateY(-950px)';
            description.style.transform = 'translateX(720px)';
            img.style.transform = 'translateX(-700px)';
        }
    });

    window.addEventListener('wheel', function (event) {
        if (event.deltaY < 0 && isSwiped) {
            isSwiped = false;
            header.style.transform = 'translateY(0)';
            contentBox.style.transform = 'translateY(0)';
            description.style.transform = 'translateX(0px)'; // Уход обратно влево
            img.style.transform = 'translateX(0px)';
        }
    });
}
isVisible();