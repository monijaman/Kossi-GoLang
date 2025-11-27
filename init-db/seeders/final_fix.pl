#!/usr/bin/perl
use strict;
use warnings;

my %devanagari_to_ascii = (
    '०' => '0', '१' => '1', '२' => '2', '३' => '3', '४' => '4',
    '५' => '5', '६' => '6', '७' => '7', '८' => '8', '९' => '9',
);

my %ascii_to_devanagari = reverse %devanagari_to_ascii;

my $in_bangla = 0;
my $in_specs = 0;

while (<>) {
    my $line = $_;
    
    # Track which section we're in
    if ($line =~ /func.*getBanglaTranslations/) {
        $in_bangla = 1;
        $in_specs = 0;
        print $line;
        next;
    } elsif ($line =~ /return map\[string\]string\{/) {
        print $line;
        next;
    } elsif ($line =~ /specs\s*:=\s*map\[string\]string\{/) {
        $in_bangla = 0;
        $in_specs = 1;
        print $line;
        next;
    } elsif ($line =~ /^\s*\}/) {
        if ($in_bangla || $in_specs) {
            $in_bangla = 0;
            $in_specs = 0;
        }
        print $line;
        next;
    }
    
    # In getBanglaTranslations map
    if ($in_bangla) {
        # Parse: "KEY": "VALUE",
        if ($line =~ /^(\s*)"([^"]*)(":\s*")([^"]*)(".*)$/) {
            my ($indent, $key, $sep1, $value, $sep2) = ($1, $2, $3, $4, $5);
            
            # Convert Devanagari back to ASCII in KEY
            foreach my $dev (keys %devanagari_to_ascii) {
                $key =~ s/$dev/$devanagari_to_ascii{$dev}/g;
            }
            
            # Convert ASCII to Devanagari in VALUE
            foreach my $digit (keys %ascii_to_devanagari) {
                $value =~ s/$digit/$ascii_to_devanagari{$digit}/g;
            }
            
            $line = $indent . '"' . $key . '"' . $sep1 . $value . $sep2 . "\n";
        }
        print $line;
        next;
    }
    
    # In specs map - convert all Devanagari to ASCII
    if ($in_specs) {
        foreach my $dev (keys %devanagari_to_ascii) {
            $line =~ s/$dev/$devanagari_to_ascii{$dev}/g;
        }
    }
    
    print $line;
}
