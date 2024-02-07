terraform {
  required_providers {
    linode = {
      source = "linode/linode"
      version = "1.16.0"
    }
  }
}

provider "linode" {
    token = "3da03191593542d87078aad7adada1c926c1e2454370ba5dbdd02993322144af"
}

resource "linode_instance" "instance-1" {
    label = "instance-1"
    image = "linode/debian10"
    region = "es-mad"
    type = "g6-standard-1"
    authorized_keys = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC+rYgNmA3yqwKrz1GnaNeKIXuQi36NUgMvAg8EZcBywCXQONQv4XXa725exLFXb8oMSBJgedGUcA1+hx65n1ZEa4tB3wBUUkF820YcBO6dqTRlHwPTs3ticbK/5OVxXBHdjRzRI87XqSVyr0/rhhH0sqxHA+WTUGjkIGsl+nmYfqqAl23ynj2qMmXgPpJvFFFmZq2pOiLZ2HKF75UDf31F+81LdVy6LtgrrkdMmd1GdBg/ITTS9rLeeAfUHeE7q5HMXhWsyXdUthk9H4nIC1fWaePIvW9+WMWXi+p/9s74J4pTjQlLPpyzNhxTvac+wYfyEYD+i6NSfCP7Hg87tj2p"]
    root_pass = "87sq*Q5!kqR3!k%o"
}

resource "linode_instance" "instance-2" {
    label = "instance-2"
    image = "linode/debian10"
    region = "es-mad"
    type = "g6-standard-1"
    authorized_keys = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC+rYgNmA3yqwKrz1GnaNeKIXuQi36NUgMvAg8EZcBywCXQONQv4XXa725exLFXb8oMSBJgedGUcA1+hx65n1ZEa4tB3wBUUkF820YcBO6dqTRlHwPTs3ticbK/5OVxXBHdjRzRI87XqSVyr0/rhhH0sqxHA+WTUGjkIGsl+nmYfqqAl23ynj2qMmXgPpJvFFFmZq2pOiLZ2HKF75UDf31F+81LdVy6LtgrrkdMmd1GdBg/ITTS9rLeeAfUHeE7q5HMXhWsyXdUthk9H4nIC1fWaePIvW9+WMWXi+p/9s74J4pTjQlLPpyzNhxTvac+wYfyEYD+i6NSfCP7Hg87tj2p"]
    root_pass = "87sq*Q5!kqR3!k%o"
}

resource "linode_instance" "instance-3" {
    label = "instance-3"
    image = "linode/debian10"
    region = "es-mad"
    type = "g6-standard-1"
    authorized_keys = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC+rYgNmA3yqwKrz1GnaNeKIXuQi36NUgMvAg8EZcBywCXQONQv4XXa725exLFXb8oMSBJgedGUcA1+hx65n1ZEa4tB3wBUUkF820YcBO6dqTRlHwPTs3ticbK/5OVxXBHdjRzRI87XqSVyr0/rhhH0sqxHA+WTUGjkIGsl+nmYfqqAl23ynj2qMmXgPpJvFFFmZq2pOiLZ2HKF75UDf31F+81LdVy6LtgrrkdMmd1GdBg/ITTS9rLeeAfUHeE7q5HMXhWsyXdUthk9H4nIC1fWaePIvW9+WMWXi+p/9s74J4pTjQlLPpyzNhxTvac+wYfyEYD+i6NSfCP7Hg87tj2p"]
    root_pass = "87sq*Q5!kqR3!k%o"
}