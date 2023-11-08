function Prompt() {
    const toast = function (c) {
      const {
        msg = "",
        position = 'top-end',
        icon = 'success'
      } = c;

      const Toast = Swal.mixin({
        toast: true,
        title: msg,
        position,
        icon,
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        didOpen: (toast) => {
          toast.addEventListener('mouseenter', Swal.stopTimer)
          toast.addEventListener('mouseleave', Swal.resumeTimer)
        }
      });

      Toast.fire({});
    };

    const success = function (c) {
      const {
        msg = '',
        title = '',
        footer = ''
      } = c;

      Swal.fire({
        icon: 'success',
        title,
        text: msg,
        footer
      })
    }

    const error = function (c) {
      const {
        msg = '',
        title = '',
        footer = ''
      } = c;

      Swal.fire({
        icon: 'error',
        title,
        text: msg,
        footer
      })
    }

    const custom = async function (c) {
      const {
        icon = "",
        msg = "",
        title = "",
        showConfirmButton = true,
        showCancelButton = true
      } = c

      const { value: result } = await Swal.fire({
        icon,
        title,
        html: msg,
        backdrop: false,
        focusConfirm: false,
        showConfirmButton,
        showCancelButton,
        willOpen: () => {
          if (c.willOpen !== undefined) {
            c.willOpen()
          }
        },
        preConfirm: () => {
          return [
            document.getElementById('start').value,
            document.getElementById('end').value
          ]
        },
        didOpen: () => {
          if (c.didOpen !== undefined) {
            c.didOpen()
          }
        }
      })

      if (result) {
        if (c.dismiss !== Swal.DismissReason.cancel) {
          if (c.callback !== undefined) {
            c.callback(result)
          } else {
            c.callback(false);
          }
        }
      } else {
        c.callback(false);
      }

    }

    return {
      toast,
      success,
      error,
      custom
    }
  }

  function setupMakeReservationButton(roomID) {
    document.getElementById('check-availability-button').addEventListener('click', function() {
        const html = `
            <form id="check-availability-form" action="" method="post" novalidate class="needs-validation">
                <div class="row">
                    <div class="col">
                        <div id="reservation-dates-modal">
                            <div class="row">
                                <div class="col">
                                    <input disabled required type="text" class="form-control" name="start" id="start" placeholder="Arrival" />
                                </div>
                                <div class="col">
                                    <input disabled required type="text" class="form-control" name="end" id="end" placeholder="Departure" />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </form>
        `; 

        attention.custom({
            title: 'Choose your dates', 
            msg: html,
            willOpen: () => {
                const elem = document.getElementById('reservation-dates-modal');
                const rp = new DateRangePicker(elem, {
                    format: 'yyyy-mm-dd',
                    buttonClass: 'btn',
                    showOnFocus: true,
                    orientation: 'top',
                    minDate: new Date()
                    })
            },
            callback: function(result) {
                
                const form = document.getElementById('check-availability-form');
                const formData = new FormData(form);
                formData.append("csrf_token","{{.CSRFToken}}")
                formData.append("room_id", roomID)

                fetch('/search-availability-json', {
                    method: 'post',
                    body: formData
                })
                    .then(response => response.json())
                    .then(data => {
                        if (data.ok) {
                            attention.custom({
                                icon: 'success',
                                msg: `<p>Room is available</p>
                                      <p>
                                            <a href="/book-room?id=${data.room_id}&s=${data.start_date}&e=${data.end_date}" 
                                                class="btn btn-primary">
                                                Book now
                                            </a>
                                      </p>`,
                                showConfirmButton: false,
                                showCancelButton: true
                            })
                        } else {
                            attention.error({
                                msg: "No Availability",
                                type: "error"
                            })
                        }
                    })
            },
            didOpen: () => {
                document.getElementById('start').removeAttribute('disabled');
                document.getElementById('end').removeAttribute('disabled');
          }
        })
      })
  }