def new_remote_control_car(nickname):
  {
    "battery_percentage": 100,
    "distance_driven_in_meters": 0,
    "nickname": nickname
  }
;

def new_remote_control_car:
  # {
  #   "battery_percentage": 100,
  #   "distance_driven_in_meters": 0,
  #   "nickname": null
  # }
  new_remote_control_car(null)
;

def display_distance:
  .distance_driven_in_meters | "\(.) meters"
;

def display_battery:
  .battery_percentage | if . == 0 then "empty"
  else "at \(.)%"
  end | "Battery \(.)"
;

def drive:
  . + if .battery_percentage > 0 then {
    "battery_percentage": (.battery_percentage - 1),
    "distance_driven_in_meters": (.distance_driven_in_meters + 20)
  }
  else .
  end
;
