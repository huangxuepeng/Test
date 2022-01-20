((doc, win) => {
    var docE = doc.documentElement;
    docE.addEventListener(
        'touchstart',
        function(event) {
            if (event.touches.length > 1) {
                console.log('缩放', JSON.stringify(event.touches));
                event.preventDefault();
            }
        },
        false
    );

    var lastTouchEnd = 0;
    docE.addEventListener(
        'touchend',
        function(event) {
            var now = Date.now();
            if (now - lastTouchEnd <= 300) {
                console.log('双击缩放');
                event.preventDefault();
            }
            lastTouchEnd = now;
        },
        false
    );

    // 阻止双指放大
    docE.addEventListener('gesturestart', function(event) {
        event.preventDefault();
    });
})(document, window);
